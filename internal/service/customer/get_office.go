package service

import (
	"context"
	"log"

	customer_models "github.com/MSFT/internal/models/customer"
	"github.com/MSFT/internal/store"
	"github.com/MSFT/internal/timestamp"
	pb "github.com/MSFT/pkg/services/customer"
)

func (s *CustomerService) GetOfficeList(ctx context.Context, in *pb.GetOfficeListRequest) (*pb.GetOfficeListResponse, error) {
	var offices []customer_models.Office
	var result []*pb.Office

	if err := store.DB.Model(&customer_models.Office{}).Find(&offices).Error; err != nil {
		log.Println("OFFICE: GetOfficeList error:\n", err)
		return nil, err
	}

	for _, item := range offices {
		result = append(result, &pb.Office{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Address:   item.Address,
			CreatedAt: timestamp.ToTimestamppb(item.CreatedAt),
		})
	}

	log.Println("OFFICE: GetOfficeList:\n", result)
	return &pb.GetOfficeListResponse{Result: result}, nil
}
