package service

import (
	"context"
	"log"

	pb "github.com/MSFT/pkg/services/customer"
)

func (s *CustomerService) GetOfficeList(ctx context.Context, in *pb.GetOfficeListRequest) (*pb.GetOfficeListResponse, error) {
	log.Println("OFFICE: get office")
	return &pb.GetOfficeListResponse{}, nil
}
