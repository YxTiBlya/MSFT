package service

import (
	"context"

	log "github.com/MSFT/internal/log"
	customer_models "github.com/MSFT/internal/models/customer"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CustomerService) GetOfficeByUUID(ctx context.Context, in *pb.GetOfficeByUUIDRequest) (*pb.GetOfficeByUUIDResponse, error) {
	var office customer_models.Office

	if err := store.DB.Model(&customer_models.Office{}).Where("uuid = ?", in.OfficeUuid).First(&office).Error; err != nil {
		log.ContextLogger.Error("GetOfficeByUUID error:", err)
		return nil, err
	}

	result := &pb.Office{
		Uuid:      office.Uuid,
		Name:      office.Name,
		Address:   office.Address,
		CreatedAt: &timestamppb.Timestamp{Seconds: office.CreatedAt.Unix()},
	}

	log.ContextLogger.Info("GetOfficeByUUID:", result)
	return &pb.GetOfficeByUUIDResponse{Result: result}, nil
}
