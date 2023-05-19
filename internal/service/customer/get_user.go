package service

import (
	"context"
	"log"

	customer_models "github.com/MSFT/internal/models/customer"
	"github.com/MSFT/internal/store"
	"github.com/MSFT/internal/timestamp"
	pb "github.com/MSFT/pkg/services/customer"
)

func (s *CustomerService) GetUserList(ctx context.Context, in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	var users []customer_models.User
	var result []*pb.User

	if err := store.DB.Model(&customer_models.User{}).Where("office_uuid = ?", in.OfficeUuid).Find(&users).Error; err != nil {
		log.Println("USER: GetUserList error:\n", err)
		return nil, err
	}

	for _, item := range users {
		result = append(result, &pb.User{
			Uuid:       item.Uuid,
			Name:       item.Name,
			OfficeUuid: item.Office_uuid,
			OfficeName: item.Office_name,
			CreatedAt:  timestamp.ToTimestamppb(item.CreatedAt),
		})
	}

	log.Println("USER: GetUserList:\n", in)
	return &pb.GetUserListResponse{Result: result}, nil
}
