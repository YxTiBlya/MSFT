package handlers

import (
	"context"
	"log"

	pb "github.com/MSFT/pkg/services/restaurant"
)

type RestaurantServer struct {
	pb.UnimplementedMenuServiceServer
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedProductServiceServer
}

func (s *RestaurantServer) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	log.Println("MENU CREATED AND ADDED TO DB")
	return &pb.CreateMenuResponse{}, nil
}

func (s *RestaurantServer) GetMenu(ctx context.Context, in *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	log.Println("MENU GET REQUEST")
	return &pb.GetMenuResponse{}, nil
}
