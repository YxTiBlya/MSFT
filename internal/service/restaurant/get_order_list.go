package service

import (
	"context"

	log "github.com/sirupsen/logrus"

	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) GetUpToDateOrderList(ctx context.Context, in *pb.GetUpToDateOrderListRequest) (*pb.GetUpToDateOrderListResponse, error) {
	log.Infoln("ORDER: get order")
	return &pb.GetUpToDateOrderListResponse{}, nil
}
