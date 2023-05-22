package service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"

	customer_models "github.com/MSFT/internal/models/customer"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/customer"
)

func (s *CustomerService) GetUserList(ctx context.Context, in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	var users []customer_models.User
	var result []*pb.User

	if err := store.DB.Model(&customer_models.User{}).Where("office_uuid = ?", in.OfficeUuid).Find(&users).Error; err != nil {
		log.Errorln("USER: GetUserList error:", err)
		return nil, err
	}

	for _, item := range users {
		result = append(result, &pb.User{
			Uuid:       item.Uuid,
			Name:       item.Name,
			OfficeUuid: item.Office_uuid,
			OfficeName: item.Office_name,
			CreatedAt:  &timestamppb.Timestamp{Seconds: item.CreatedAt.Unix()},
		})
	}

	log.Infoln("USER: GetUserList:", in)
	return &pb.GetUserListResponse{Result: result}, nil
}
