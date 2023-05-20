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

func (s *CustomerService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var office customer_models.Office

	store.DB.Model(&customer_models.Office{}).Where("uuid = ?", in.OfficeUuid).First(&office)

	created_at_time := timestamppb.New(time.Now())
	user := customer_models.User{
		Name:        in.Name,
		Office_uuid: office.Uuid,
		Office_name: office.Name,
		CreatedAt:   fmt.Sprintf("%v.%v", created_at_time.Seconds, created_at_time.Nanos),
	}

	if err := store.DB.Model(&customer_models.User{}).Create(&user).Error; err != nil {
		log.Println("USER: CreateUser error:\n", err)
		return nil, err
	}

	log.Println("USER: CreateUser:\n", in)
	return &pb.CreateUserResponse{}, nil
}