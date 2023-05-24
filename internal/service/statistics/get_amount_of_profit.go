package service

import (
	"context"
	"time"

	log "github.com/MSFT/internal/log"
	statistics_models "github.com/MSFT/internal/models/statistics"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/statistics"
)

func (s *StatisticsService) GetAmountOfProfit(ctx context.Context, in *pb.GetAmountOfProfitRequest) (*pb.GetAmountOfProfitResponse, error) {
	var statistics []statistics_models.Statistics

	startDate := time.Unix(in.StartDate.Seconds, int64(in.StartDate.Nanos))
	endDate := time.Unix(in.EndDate.Seconds, int64(in.EndDate.Nanos))

	if err := store.DB.Model(&statistics_models.Statistics{}).Where("created_at >= ? AND created_at <= ?", startDate, endDate).Find(&statistics).Error; err != nil {
		log.ContextLogger.Error("GetAmountOfProfit error:", err)
		return nil, err
	}

	var profit float64 = 0
	for _, item := range statistics {
		profit += item.Profit
	}

	log.ContextLogger.Info("GetAmountOfProfit:", profit)
	return &pb.GetAmountOfProfitResponse{Profit: profit}, nil
}
