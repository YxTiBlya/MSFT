package service

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) GetUpToDateOrderList(ctx context.Context, in *pb.GetUpToDateOrderListRequest) (*pb.GetUpToDateOrderListResponse, error) {
	var orders restaurant_models.Orders

	nowTime := time.Now()
	startTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
	endTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 23, 59, 59, 0, time.Local)

	if err := store.DB.Model(&restaurant_models.Orders{}).Where("created_at >= ? AND created_at <= ?", startTime, endTime).First(&orders).Error; err != nil {
		log.Errorln("ORDER: GetUpToDateOrderList error:", err)
		return nil, err
	}

	result := &pb.GetUpToDateOrderListResponse{
		TotalOrders:          orders.TotalOrders.ToGRPCModel(),
		TotalOrdersByCompany: orders.TotalOrdersByCompany.ToGRPCModel(),
	}
	log.Infoln("ORDER: GetUpToDateOrderList:", result)
	return result, nil
}
