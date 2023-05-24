package service

import (
	"context"
	"sort"
	"time"

	statistics_models "github.com/MSFT/internal/models/statistics"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/statistics"
	log "github.com/sirupsen/logrus"
)

func (s *StatisticsService) TopProducts(ctx context.Context, in *pb.TopProductsRequest) (*pb.TopProductsResponse, error) {
	var statistics []statistics_models.Statistics

	startDate := time.Unix(in.StartDate.Seconds, int64(in.StartDate.Nanos))
	endDate := time.Unix(in.EndDate.Seconds, int64(in.EndDate.Nanos))

	if err := store.DB.Model(&statistics_models.Statistics{}).Where("created_at >= ? AND created_at <= ?", startDate, endDate).Find(&statistics).Error; err != nil {
		log.Errorln("STATISTICS: TopProducts error:", err)
		return nil, err
	}

	result := make([]*pb.Product, 0, len(statistics)*5)
	for _, statistic := range statistics {
		result = append(result, statistic.TopProducts.ToGRPCModel()...)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	log.Infoln("STATISTICS: TopProducts:", result)
	return &pb.TopProductsResponse{Result: result}, nil
}
