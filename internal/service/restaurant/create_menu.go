package service

import (
	"context"
	"log"
	"time"

	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
)

func (s *RestaurantService) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	var salads, garnishes, meats, soups, drinks, desserts restaurant_models.ProductArray

	store.DB.Model(&restaurant_models.Product{}).Where("uuid IN ?", in.Salads).Find(&salads)
	store.DB.Model(&restaurant_models.Product{}).Where("uuid IN ?", in.Garnishes).Find(&garnishes)
	store.DB.Model(&restaurant_models.Product{}).Where("uuid IN ?", in.Meats).Find(&meats)
	store.DB.Model(&restaurant_models.Product{}).Where("uuid IN ?", in.Soups).Find(&soups)
	store.DB.Model(&restaurant_models.Product{}).Where("uuid IN ?", in.Drinks).Find(&drinks)
	store.DB.Model(&restaurant_models.Product{}).Where("uuid IN ?", in.Desserts).Find(&desserts)

	menu := restaurant_models.Menu{
		OnDate:          in.OnDate.AsTime(),
		OpeningRecordAt: in.OpeningRecordAt.AsTime(),
		ClosingRecordAt: in.ClosingRecordAt.AsTime(),
		Salads:          salads,
		Garnishes:       garnishes,
		Meats:           meats,
		Soups:           soups,
		Drinks:          drinks,
		Desserts:        desserts,
		CreatedAt:       time.Now(),
	}

	if err := store.DB.Model(&restaurant_models.Menu{}).Create(&menu).Error; err != nil {
		log.Println("MENU: CreateMenu error:\n", err)
		return nil, err
	}

	log.Println("MENU: CreateMenu:\n", in)
	return &pb.CreateMenuResponse{}, nil
}
