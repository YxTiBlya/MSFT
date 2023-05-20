package service

import (
	"context"
	"log"
	"time"

	customer_models "github.com/MSFT/internal/models/customer"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/customer"
)

func (s *CustomerService) CreateOffice(ctx context.Context, in *pb.CreateOfficeRequest) (*pb.CreateOfficeResponse, error) {
	office := customer_models.Office{
		Name:      in.Name,
		Address:   in.Address,
		CreatedAt: time.Now(),
	}

	if err := store.DB.Model(&customer_models.Office{}).Create(&office).Error; err != nil {
		log.Println("OFFICE: CreateOffice error:\n", err)
		return nil, err
	}

	log.Println("OFFICE: CreateOffice:\n", in)
	return &pb.CreateOfficeResponse{}, nil
}
