package service

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"

	customer_models "github.com/MSFT/internal/models/customer"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/customer"
)

func (s *CustomerService) GetOfficeList(ctx context.Context, in *pb.GetOfficeListRequest) (*pb.GetOfficeListResponse, error) {
	var offices []customer_models.Office
	var result []*pb.Office

	if err := store.DB.Model(&customer_models.Office{}).Find(&offices).Error; err != nil {
		log.Errorln("OFFICE: GetOfficeList error:", err)
		return nil, err
	}

	for _, item := range offices {
		result = append(result, &pb.Office{
			Uuid:      item.Uuid,
			Name:      item.Name,
			Address:   item.Address,
			CreatedAt: &timestamppb.Timestamp{Seconds: item.CreatedAt.Unix()},
		})
	}

	log.Infoln("OFFICE: GetOfficeList:", result)
	return &pb.GetOfficeListResponse{Result: result}, nil
}
