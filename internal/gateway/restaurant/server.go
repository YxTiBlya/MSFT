package gateway_restaurant

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	log "github.com/MSFT/internal/log"

	"github.com/MSFT/internal/cfg"
	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/internal/rabbitmq"
	service "github.com/MSFT/internal/service/restaurant"
	"github.com/MSFT/internal/store"
	"github.com/MSFT/pkg/services/customer"
	pb "github.com/MSFT/pkg/services/restaurant"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RestaurantServer struct{}

func (rs *RestaurantServer) RunGRPCServer(cfg *cfg.Config, s *grpc.Server) {
	service := &service.RestaurantService{}
	pb.RegisterMenuServiceServer(s, service)
	pb.RegisterOrderServiceServer(s, service)
	pb.RegisterProductServiceServer(s, service)

	l, err := net.Listen("tcp", fmt.Sprintf("%v:%d", cfg.General_host, cfg.Restaurant_grpc_service_port))
	if err != nil {
		log.ContextLogger.Fatal("failed to listen:", err.Error())
	}

	log.ContextLogger.Infof("starting listening grpc server at %v", fmt.Sprintf("%v:%d", cfg.General_host, cfg.Restaurant_grpc_service_port))
	if err := s.Serve(l); err != nil {
		log.ContextLogger.Fatal("error service grpc server:", err.Error())
	}
}

func (rs *RestaurantServer) RuntHTTPServer(ctx context.Context, cfg *cfg.Config, mux *runtime.ServeMux) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := fmt.Sprintf("%v:%d", cfg.General_host, cfg.Restaurant_grpc_service_port)

	if err := pb.RegisterMenuServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.ContextLogger.Fatal(err)
	}

	if err := pb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.ContextLogger.Fatal(err)
	}

	if err := pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.ContextLogger.Fatal(err)
	}

	log.ContextLogger.Infof("starting listening http server at %s", fmt.Sprintf("%v:%d", cfg.General_host, cfg.Restaurant_http_service_port))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Restaurant_http_service_port), mux); err != nil {
		log.ContextLogger.Fatal("error service http server:", err.Error())
	}
}

func (rs *RestaurantServer) RunRabbitMQReciever(cfg *cfg.Config) {
	msgs, err := rabbitmq.RecieveOrder(cfg)
	if err != nil {
		log.ContextLogger.Fatalln("failed to recieve at rabbitmq:", err)
	}

	var forever chan struct{}
	log.ContextLogger.Infoln("starting listening order queue at rabbitmq")
	go func() {
		for d := range msgs {
			log.ContextLogger.Infof("recieved message: %s", d.Body)

			orderRequest := customer.CreateOrderRequest{}
			json.Unmarshal(d.Body, &orderRequest)

			var orders restaurant_models.Orders
			var findedOrder bool = true

			nowTime := time.Now()
			startTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
			endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 23, 59, 59, 0, time.Local)

			if err := store.DB.Model(&restaurant_models.Orders{}).Where("created_at >= ? AND created_at <= ?", startTime, endTime).First(&orders).Error; err != nil {
				log.ContextLogger.Infoln("not found TotalOrders of today")
				findedOrder = false
			}

			if err := service.UpdateOrderList(&orders, &orderRequest); err != nil {
				log.ContextLogger.Errorln("error to update total order list:", err)
				continue
			}

			if findedOrder {
				if err := store.DB.Model(&restaurant_models.Orders{}).Omit("created_at").Where("id = ?", orders.Id).Updates(&orders).Error; err != nil {
					log.ContextLogger.Errorln("failed to update order:", err)
					continue
				}
			} else {
				orders.CreatedAt = nowTime
				if err := store.DB.Model(&restaurant_models.Orders{}).Create(&orders).Error; err != nil {
					log.ContextLogger.Errorln("failed to create order:", err)
					continue
				}
			}

			log.ContextLogger.Infoln("updated total order:", orders)
		}
	}()

	<-forever
}
