package service

import (
	"context"
	"time"

	statistics_models "github.com/MSFT/internal/models/statistics"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/statistics"
	log "github.com/sirupsen/logrus"
)

func (s *StatisticsService) GetAmountOfProfit(ctx context.Context, in *pb.GetAmountOfProfitRequest) (*pb.GetAmountOfProfitResponse, error) {
	var statistics []statistics_models.Statistics

	startDate := time.Unix(in.StartDate.Seconds, int64(in.StartDate.Nanos))
	endDate := time.Unix(in.EndDate.Seconds, int64(in.EndDate.Nanos))

	if err := store.DB.Model(&statistics_models.Statistics{}).Where("created_at >= ? AND created_at <= ?", startDate, endDate).Find(&statistics).Error; err != nil {
		log.Errorln("STATISTICS: GetAmountOfProfit error:", err)
		return nil, err
	}

	var profit float64 = 0
	for _, item := range statistics {
		profit += item.Profit
	}

	log.Infoln("STATISTICS: GetAmountOfProfit:", profit)
	return &pb.GetAmountOfProfitResponse{Profit: profit}, nil
}
