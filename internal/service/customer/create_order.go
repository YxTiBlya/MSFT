package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/rabbitmq"
	pb "github.com/MSFT/pkg/services/customer"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (s *CustomerService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
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
		log.Println("error publish the order:\n", err.Error())
	}

	return &pb.CreateOrderResponse{}, nil
}
