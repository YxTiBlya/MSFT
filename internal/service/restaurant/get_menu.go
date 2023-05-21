package service

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/internal/store"
	"github.com/MSFT/internal/timestamp"
	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) GetMenu(ctx context.Context, in *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	var menu restaurant_models.Menu

	onDate := time.Unix(in.OnDate.Seconds, int64(in.OnDate.Nanos))
	onDateY := onDate.Year()
	onDateM := onDate.Month()
	onDateD := onDate.Day()

	startDate := time.Date(onDateY, onDateM, onDateD, 0, 0, 0, 0, time.Local)
	endDate := time.Date(onDateY, onDateM, onDateD, 23, 59, 59, 0, time.Local)

	if err := store.DB.Model(&restaurant_models.Menu{}).Where("on_date >= ? AND on_date <= ?", startDate, endDate).First(&menu).Error; err != nil {
		log.Errorln("MENU: GetMenu error:", err.Error())
		return nil, err
	}

	result := &pb.Menu{
		Uuid:            menu.Uuid,
		OnDate:          timestamp.ToTimestamppb(menu.OnDate),
		OpeningRecordAt: timestamp.ToTimestamppb(menu.OpeningRecordAt),
		ClosingRecordAt: timestamp.ToTimestamppb(menu.ClosingRecordAt),
		Salads:          menu.Salads.ToGRPCModel(),
		Garnishes:       menu.Garnishes.ToGRPCModel(),
		Meats:           menu.Meats.ToGRPCModel(),
		Soups:           menu.Soups.ToGRPCModel(),
		Drinks:          menu.Drinks.ToGRPCModel(),
		Desserts:        menu.Desserts.ToGRPCModel(),
		CreatedAt:       timestamp.ToTimestamppb(menu.CreatedAt),
	}

	log.Infoln("MENU: GetMenu:", result)
	return &pb.GetMenuResponse{Menu: result}, nil
}
