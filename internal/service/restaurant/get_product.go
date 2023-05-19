package service

import (
	"context"
	"log"

	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/internal/store"
	"github.com/MSFT/internal/timestamp"
	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) GetProduct(ctx context.Context, in *pb.GetProductListRequest) (*pb.GetProductListResponse, error) {
	var products []restaurant_models.Product
	var result []*pb.Product

	if err := store.DB.Model(&restaurant_models.Product{}).Find(&products).Error; err != nil {
		log.Println("PRODUCT: GetProduct error:\n", err)
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
			CreatedAt:   timestamp.ToTimestamppb(item.CreatedAt),
		})
	}

	log.Println("PRODUCT: GetProduct:\n", result)
	return &pb.GetProductListResponse{Result: result}, nil
}
