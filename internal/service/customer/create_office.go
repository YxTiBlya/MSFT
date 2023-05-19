package service

import (
	"context"
	"fmt"
	"log"
	"time"

	customer_models "github.com/MSFT/internal/models/customer"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CustomerService) CreateOffice(ctx context.Context, in *pb.CreateOfficeRequest) (*pb.CreateOfficeResponse, error) {
	created_at_time := timestamppb.New(time.Now())
	office := customer_models.Office{
		Name:      in.Name,
		Address:   in.Address,
		CreatedAt: fmt.Sprintf("%v.%v", created_at_time.Seconds, created_at_time.Nanos),
	}

	if err := store.DB.Model(&customer_models.Office{}).Create(&office).Error; err != nil {
		log.Println("OFFICE: CreateOffice error:\n", err)
		return nil, err
	}

	log.Println("OFFICE: CreateOffice:\n", in)
	return &pb.CreateOfficeResponse{}, nil
}
