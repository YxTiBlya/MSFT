package service

import (
	"context"

	log "github.com/MSFT/internal/log"
	"google.golang.org/protobuf/types/known/timestamppb"

	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) GetProductByUUID(ctx context.Context, in *pb.GetProductByUUIDRequest) (*pb.GetProductByUUIDResponse, error) {
	var product restaurant_models.Product

	if err := store.DB.Model(&restaurant_models.Product{}).Where("uuid = ?", in.ProductUuid).First(&product).Error; err != nil {
		log.ContextLogger.Error("GetProductByUUID error:", err)
		return nil, err
	}

	result := &pb.Product{
		Uuid:        product.Uuid,
		Name:        product.Name,
		Description: product.Description,
		Type:        pb.ProductType(product.Type),
		Weight:      product.Weight,
		Price:       product.Price,
		CreatedAt:   &timestamppb.Timestamp{Seconds: product.CreatedAt.Unix()},
	}

	log.ContextLogger.Info("GetProductByUUID:", result)
	return &pb.GetProductByUUIDResponse{Result: result}, nil
}
