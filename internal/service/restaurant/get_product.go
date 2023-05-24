package service

import (
	"context"

	log "github.com/MSFT/internal/log"
	"google.golang.org/protobuf/types/known/timestamppb"

	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) GetProduct(ctx context.Context, in *pb.GetProductListRequest) (*pb.GetProductListResponse, error) {
	var products []restaurant_models.Product
	var result []*pb.Product

	if err := store.DB.Model(&restaurant_models.Product{}).Find(&products).Error; err != nil {
		log.ContextLogger.Error("GetProduct error:", err)
		return nil, err
	}

	for _, item := range products {
		result = append(result, &pb.Product{
			Uuid:        item.Uuid,
			Name:        item.Name,
			Description: item.Description,
			Type:        pb.ProductType(item.Type),
			Weight:      item.Weight,
			Price:       item.Price,
			CreatedAt:   &timestamppb.Timestamp{Seconds: item.CreatedAt.Unix()},
		})
	}

	log.ContextLogger.Info("GetProduct:", result)
	return &pb.GetProductListResponse{Result: result}, nil
}
