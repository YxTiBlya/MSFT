package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MSFT/internal/models"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *RestaurantService) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	var salads, garnishes, meats, soups, drinks, desserts models.ProductArray

	store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Salads).Find(&salads)
	store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Garnishes).Find(&garnishes)
	store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Meats).Find(&meats)
	store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Soups).Find(&soups)
	store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Drinks).Find(&drinks)
	store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Desserts).Find(&desserts)

	created_at_time := timestamppb.New(time.Now())
	menu := models.Menu{
		OnDate:          fmt.Sprintf("%v.%v", in.OnDate.Seconds, in.OnDate.Nanos),
		OpeningRecordAt: fmt.Sprintf("%v.%v", in.OpeningRecordAt.Seconds, in.OpeningRecordAt.Nanos),
		ClosingRecordAt: fmt.Sprintf("%v.%v", in.ClosingRecordAt.Seconds, in.ClosingRecordAt.Nanos),
		Salads:          salads,
		Garnishes:       garnishes,
		Meats:           meats,
		Soups:           soups,
		Drinks:          drinks,
		Desserts:        desserts,
		CreatedAt:       fmt.Sprintf("%v.%v", created_at_time.Seconds, created_at_time.Nanos),
	}

	if err := store.DB.Model(&models.Menu{}).Create(&menu).Error; err != nil {
		log.Println("MENU: CreateMenu error:\n", err)
		return nil, err
	}

	log.Println("MENU: CreateMenu:\n", in)
	return &pb.CreateMenuResponse{}, nil
}
