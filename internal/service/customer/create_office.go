package service

import (
	"context"
	"log"

	pb "github.com/MSFT/pkg/services/customer"
)

func (s *CustomerService) CreateOffice(ctx context.Context, in *pb.CreateOfficeRequest) (*pb.CreateOfficeResponse, error) {
	log.Println("OFFICE: created office:\n", in)
	return &pb.CreateOfficeResponse{}, nil
}
