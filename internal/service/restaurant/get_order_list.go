package service

import (
	"context"
	"log"

	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) GetUpToDateOrderList(ctx context.Context, in *pb.GetUpToDateOrderListRequest) (*pb.GetUpToDateOrderListResponse, error) {
	log.Println("ORDER: get order")
	return &pb.GetUpToDateOrderListResponse{}, nil
}
