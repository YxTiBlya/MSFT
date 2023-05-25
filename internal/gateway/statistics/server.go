package gateway_statistics

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	log "github.com/MSFT/internal/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/MSFT/internal/cfg"
	statistics_models "github.com/MSFT/internal/models/statistics"
	"github.com/MSFT/internal/rabbitmq"
	service "github.com/MSFT/internal/service/statistics"
	"github.com/MSFT/internal/store"
	"github.com/MSFT/pkg/services/customer"
	pb "github.com/MSFT/pkg/services/statistics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type StatisticsServer struct{}

func (ss *StatisticsServer) RunGRPCServer(cfg *cfg.Config, s *grpc.Server) {
	service := &service.StatisticsService{}
	pb.RegisterStatisticsServiceServer(s, service)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Statistics_grpc_service_port))
	if err != nil {
		log.ContextLogger.Fatal("failed to listen:", err.Error())
	}

	log.ContextLogger.Infof("starting listening grpc server at %v", l.Addr())
	if err := s.Serve(l); err != nil {
		log.ContextLogger.Fatal("error service grpc server:", err.Error())
	}
}

func (ss *StatisticsServer) RuntHTTPServer(ctx context.Context, cfg *cfg.Config, mux *runtime.ServeMux) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := fmt.Sprintf("localhost:%d", cfg.Statistics_grpc_service_port)

	if err := pb.RegisterStatisticsServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.ContextLogger.Fatal(err)
	}

	log.ContextLogger.Infof("starting listening http server at %s", fmt.Sprintf("localhost:%d", cfg.Statistics_http_service_port))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Statistics_http_service_port), mux); err != nil {
		log.ContextLogger.Fatal("error service http server:", err.Error())
	}
}

func (ss *StatisticsServer) RunRabbitMQReciever(cfg *cfg.Config) {
	msgs, err := rabbitmq.RecieveOrder(cfg)
	if err != nil {
		log.ContextLogger.Fatal("failed to recieve at rabbitmq:", err)
	}

	var forever chan struct{}
	log.ContextLogger.Info("starting listening order queue at rabbitmq")
	go func() {
		for d := range msgs {
			log.ContextLogger.Infof("recieved message: %s", d.Body)

			orderRequest := customer.CreateOrderRequest{}
			json.Unmarshal(d.Body, &orderRequest)

			var statistics statistics_models.Statistics
			var findedOrder bool = true

			nowTime := time.Now()
			startTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
			endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 23, 59, 59, 0, time.Local)

			if err := store.DB.Model(&statistics_models.Statistics{}).Where("created_at >= ? AND created_at <= ?", startTime, endTime).First(&statistics).Error; err != nil {
				log.ContextLogger.Info("not found statistics of today")
				findedOrder = false
			}

			if err := service.UpdateStatisticsList(&statistics, &orderRequest); err != nil {
				log.ContextLogger.Error("error to update statistics list:", err)
				continue
			}

			if findedOrder {
				if err := store.DB.Model(&statistics_models.Statistics{}).Omit("created_at").Where("id = ?", statistics.Id).Updates(&statistics).Error; err != nil {
					log.ContextLogger.Error("failed to update statistics:", err)
					continue
				}
			} else {
				statistics.CreatedAt = nowTime
				if err := store.DB.Model(&statistics_models.Statistics{}).Create(&statistics).Error; err != nil {
					log.ContextLogger.Error("failed to create statistics record:", err)
					continue
				}
			}

			log.ContextLogger.Info("updated statistics:", statistics)
		}
	}()

	<-forever
}
