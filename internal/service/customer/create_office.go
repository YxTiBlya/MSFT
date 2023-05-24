package service

import (
	"context"
	"time"

	log "github.com/MSFT/internal/log"

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
		log.ContextLogger.Error("CreateOffice error:", err)
		return nil, err
	}

	log.ContextLogger.Info("CreateOffice:", in)
	return &pb.CreateOfficeResponse{}, nil
}
