package service

import (
	"context"

	log "github.com/MSFT/internal/log"
	customer_models "github.com/MSFT/internal/models/customer"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CustomerService) GetUserByUUID(ctx context.Context, in *pb.GetUserByUUIDRequest) (*pb.GetUserByUUIDResponse, error) {
	var user customer_models.User

	if err := store.DB.Model(&customer_models.User{}).Where("uuid = ?", in.UserUuid).First(&user).Error; err != nil {
		log.ContextLogger.Error("GetUserByUUID error:", err)
		return nil, err
	}

	result := &pb.User{
		Uuid:       user.Uuid,
		Name:       user.Name,
		OfficeUuid: user.Office_uuid,
		OfficeName: user.Office_name,
		CreatedAt:  &timestamppb.Timestamp{Seconds: user.CreatedAt.Unix()},
	}

	log.ContextLogger.Info("GetUserByUUID:", result)
	return &pb.GetUserByUUIDResponse{Result: result}, nil
}
