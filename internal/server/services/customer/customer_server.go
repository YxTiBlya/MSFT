package customer_handlers

import (
	"context"
	"log"

	pb "github.com/MSFT/pkg/services/customer"
)

type CustomerServer struct {
	pb.UnimplementedOfficeServiceServer
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedUserServiceServer
}

func (s *CustomerServer) CreateOffice(ctx context.Context, in *pb.CreateOfficeRequest) (*pb.CreateOfficeResponse, error) {
	log.Println("OFFICE: created office:\n", in)
	return &pb.CreateOfficeResponse{}, nil
}

func (s *CustomerServer) GetOfficeList(ctx context.Context, in *pb.GetOfficeListRequest) (*pb.GetOfficeListResponse, error) {
	log.Println("OFFICE: get office")
	return &pb.GetOfficeListResponse{}, nil
}
