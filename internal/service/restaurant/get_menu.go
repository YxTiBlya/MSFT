package service

import (
	"context"
	"log"

	"github.com/MSFT/internal/models"
	"github.com/MSFT/internal/store"
	"github.com/MSFT/internal/timestamp"
	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) GetMenu(ctx context.Context, in *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	var menu models.Menu
	if err := store.DB.Model(&models.Menu{}).First(&menu).Error; err != nil {
		log.Println("MENU: GetMenu error:\n", err.Error())
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

	log.Println("MENU: GetMenu:\n", result)
	return &pb.GetMenuResponse{Menu: result}, nil
}
