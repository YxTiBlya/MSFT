package service

import (
	"context"
	"fmt"
	"log"

	"github.com/MSFT/internal/cfg"
	pb "github.com/MSFT/pkg/services/customer"
	"github.com/MSFT/pkg/services/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CustomerService) GetActualMenu(ctx context.Context, in *pb.GetActualMenuRequest) (*pb.GetActualMenuResponse, error) {
	config := cfg.GetConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", config.General_host, config.Restaurant_grpc_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("ORDER: GetActualMenu error:\n", err.Error())
		return nil, err
	}
	defer conn.Close()

	client := restaurant.NewMenuServiceClient(conn)
	response, err := client.GetMenu(context.Background(), &restaurant.GetMenuRequest{OnDate: timestamppb.Now()})
	if err != nil {
		log.Println("ORDER: GetActualMenu error:\n", err.Error())
		return nil, err
	}

	actualMenu := &pb.GetActualMenuResponse{
		Salads:    toCustomerProduct(response.Menu.Salads),
		Garnishes: toCustomerProduct(response.Menu.Garnishes),
		Meats:     toCustomerProduct(response.Menu.Meats),
		Soups:     toCustomerProduct(response.Menu.Soups),
		Drinks:    toCustomerProduct(response.Menu.Drinks),
		Desserts:  toCustomerProduct(response.Menu.Desserts),
	}

	log.Println("ORDER: GetActualMenu:\n", actualMenu)
	return actualMenu, nil
}

func toCustomerProduct(products []*restaurant.Product) []*pb.Product {
	customerProducts := make([]*pb.Product, 0, len(products))

	for _, product := range products {
		customerProducts = append(customerProducts, &pb.Product{
			Uuid:        product.Uuid,
			Name:        product.Name,
			Description: product.Description,
			Type:        pb.CustomerProductType(product.Type),
			Weight:      product.Weight,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt,
		})
	}

	return customerProducts
}
