package service

import (
	pb "github.com/MSFT/pkg/services/customer"
)

type CustomerService struct {
	pb.UnimplementedOfficeServiceServer
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedUserServiceServer
}
