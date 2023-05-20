package service

import (
	"context"
	"log"
	"time"

	customer_models "github.com/MSFT/internal/models/customer"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/customer"
)

func (s *CustomerService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var office customer_models.Office

	store.DB.Model(&customer_models.Office{}).Where("uuid = ?", in.OfficeUuid).First(&office)

	user := customer_models.User{
		Name:        in.Name,
		Office_uuid: office.Uuid,
		Office_name: office.Name,
		CreatedAt:   time.Now(),
	}

	if err := store.DB.Model(&customer_models.User{}).Create(&user).Error; err != nil {
		log.Println("USER: CreateUser error:\n", err)
		return nil, err
	}

	log.Println("USER: CreateUser:\n", in)
	return &pb.CreateUserResponse{}, nil
}
