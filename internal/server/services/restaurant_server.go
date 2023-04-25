package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/MSFT/internal/models"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	time := timestamppb.New(time.Now())
	product := models.Products{
		Name:        in.Name,
		Description: in.Description,
		Type:        pb.ProductType_value[in.Type.String()],
		Weight:      in.Weight,
		Price:       in.Price,
		CreatedAt:   fmt.Sprintf("%v.%v", time.Seconds, time.Nanos),
	}

	if err := store.DB.Model(&models.Products{}).Create(&product).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateProductResponse{}, nil
}

func (s *RestaurantServer) GetProduct(ctx context.Context, in *pb.GetProductListRequest) (*pb.GetProductListResponse, error) {
	log.Println("PRODUCT: get products")

	var products []models.Products
	var result []*pb.Product

	if err := store.DB.Model(&models.Products{}).Find(&products).Error; err != nil {
		return nil, err
	}
	log.Println(products)

	for _, item := range products {
		created_at_string := strings.Split(item.CreatedAt, ".")
		secs, _ := strconv.Atoi(created_at_string[0])
		nans, _ := strconv.Atoi(created_at_string[1])

		result = append(result, &pb.Product{
			Uuid:        item.Id,
			Name:        item.Name,
			Description: item.Description,
			Type:        pb.ProductType(item.Type),
			Weight:      item.Weight,
			Price:       item.Price,
			CreatedAt:   &timestamppb.Timestamp{Seconds: int64(secs), Nanos: int32(nans)},
		})
	}

	return &pb.GetProductListResponse{Result: result}, nil
}

func (s *RestaurantServer) GetUpToDateOrderList(ctx context.Context, in *pb.GetUpToDateOrderListRequest) (*pb.GetUpToDateOrderListResponse, error) {
	log.Println("ORDER: get order")
	return &pb.GetUpToDateOrderListResponse{}, nil
}
