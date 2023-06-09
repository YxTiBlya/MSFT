package service

import (
	"context"
	"time"

	log "github.com/MSFT/internal/log"

	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := restaurant_models.Product{
		Name:        in.Name,
		Description: in.Description,
		Type:        pb.ProductType_value[in.Type.String()],
		Weight:      in.Weight,
		Price:       in.Price,
		CreatedAt:   time.Now(),
	}

	if err := store.DB.Model(&restaurant_models.Product{}).Create(&product).Error; err != nil {
		log.ContextLogger.Error("CreateProduct error:", err)
		return nil, err
	}

	log.ContextLogger.Info("CreateProduct:", in)
	return &pb.CreateProductResponse{}, nil
}
