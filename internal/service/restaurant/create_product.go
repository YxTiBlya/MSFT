package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MSFT/internal/models"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *RestaurantService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	nowTime := timestamppb.New(time.Now())

	product := models.Product{
		Name:        in.Name,
		Description: in.Description,
		Type:        pb.ProductType_value[in.Type.String()],
		Weight:      in.Weight,
		Price:       in.Price,
		CreatedAt:   fmt.Sprintf("%v.%v", nowTime.Seconds, nowTime.Nanos),
	}

	if err := store.DB.Model(&models.Product{}).Create(&product).Error; err != nil {
		log.Println("PRODUCT: CreateProduct error:\n", err)
		return nil, err
	}

	log.Println("PRODUCT: CreateProduct:\n", in)
	return &pb.CreateProductResponse{}, nil
}
