package service

import (
	pb "github.com/MSFT/pkg/services/restaurant"
)

type RestaurantService struct {
	pb.UnimplementedMenuServiceServer
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedProductServiceServer
}
