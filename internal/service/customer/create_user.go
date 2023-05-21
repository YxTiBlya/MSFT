package service

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

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
		log.Errorln("USER: CreateUser error:", err)
		return nil, err
	}

	log.Infoln("USER: CreateUser:", in)
	return &pb.CreateUserResponse{}, nil
}
