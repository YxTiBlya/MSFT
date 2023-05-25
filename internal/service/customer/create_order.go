package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	log "github.com/MSFT/internal/log"

	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/rabbitmq"
	pb "github.com/MSFT/pkg/services/customer"
	"github.com/MSFT/pkg/services/restaurant"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CustomerService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	config := cfg.GetConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", config.Restaurant_host, config.Restaurant_grpc_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.ContextLogger.Error("CreateOrder error:", err.Error())
		return nil, err
	}
	defer conn.Close()

	client := restaurant.NewMenuServiceClient(conn)
	response, err := client.GetMenu(context.Background(), &restaurant.GetMenuRequest{OnDate: timestamppb.Now()})
	if err != nil {
		log.ContextLogger.Error("CreateOrder error:", err.Error())
		return nil, err
	}

	openingRecord := time.Unix(response.Menu.OpeningRecordAt.Seconds, int64(response.Menu.OpeningRecordAt.Nanos))
	closingRecord := time.Unix(response.Menu.ClosingRecordAt.Seconds, int64(response.Menu.ClosingRecordAt.Nanos))
	nowTime := time.Now()
	if nowTime.After(openingRecord) && nowTime.Before(closingRecord) {
		lctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		msg, _ := json.Marshal(in)
		if err := rabbitmq.Chann.PublishWithContext(lctx,
			cfg.GetConfig().Rabbitmq_queue_name,
			"",
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        msg,
			},
		); err != nil {
			log.ContextLogger.Error("error publish the order:", err.Error())
		}

		log.ContextLogger.Info("CreateOrder sended msg:", in)
		return &pb.CreateOrderResponse{}, nil
	}

	log.ContextLogger.Error("CreateOrder error: the time for orders has passed")
	return nil, errors.New("the time for orders has passed")
}
