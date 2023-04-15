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
	log.Println("MENU: created menu:\n", in)
	return &pb.CreateMenuResponse{}, nil
}

func (s *RestaurantServer) GetMenu(ctx context.Context, in *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	log.Println("MENU: get menu")
	return &pb.GetMenuResponse{}, nil
}

func (s *RestaurantServer) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	log.Println("PRODUCT: create product:\n", in)
	return &pb.CreateProductResponse{}, nil
}

func (s *RestaurantServer) GetProduct(ctx context.Context, in *pb.GetProductListRequest) (*pb.GetProductListResponse, error) {
	log.Println("PRODUCT: get products")
	return &pb.GetProductListResponse{}, nil
}

func (s *RestaurantServer) GetUpToDateOrderList(ctx context.Context, in *pb.GetUpToDateOrderListRequest) (*pb.GetUpToDateOrderListResponse, error) {
	log.Println("ORDER: get order")
	return &pb.GetUpToDateOrderListResponse{}, nil
}
