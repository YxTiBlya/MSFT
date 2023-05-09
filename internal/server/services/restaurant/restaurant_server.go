package restaurant_handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/MSFT/internal/models"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RestaurantServer struct {
	pb.UnimplementedMenuServiceServer
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedProductServiceServer
}

func (s *RestaurantServer) CreateMenu(ctx context.Context, in *pb.CreateMenuRequest) (*pb.CreateMenuResponse, error) {
	log.Println("MENU: created menu:\n", in)

	var salads, garnishes, meats, soups, drinks, desserts models.ProductArray

	// :DDDD
	if err := store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Salads).Find(&salads).Error; err != nil {
		log.Println(err)
	}
	if err := store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Garnishes).Find(&garnishes).Error; err != nil {
		log.Println(err)
	}
	if err := store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Meats).Find(&meats).Error; err != nil {
		log.Println(err)
	}
	if err := store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Soups).Find(&soups).Error; err != nil {
		log.Println(err)
	}
	if err := store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Drinks).Find(&drinks).Error; err != nil {
		log.Println(err)
	}
	if err := store.DB.Model(&models.Product{}).Where("uuid IN ?", in.Desserts).Find(&desserts).Error; err != nil {
		log.Println(err)
	}

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
		log.Println(err)
		return nil, err
	}

	return &pb.CreateMenuResponse{}, nil
}

func (s *RestaurantServer) GetMenu(ctx context.Context, in *pb.GetMenuRequest) (*pb.GetMenuResponse, error) {
	log.Println("MENU: get menu")

	var menu models.Menu

	if err := store.DB.Model(&models.Menu{}).First(&menu).Error; err != nil {
		return nil, err
	}

	onDateTMP := strings.Split(menu.OnDate, ".")
	onDateTMPSecs, _ := strconv.Atoi(onDateTMP[0])
	onDateTMPNSecs, _ := strconv.Atoi(onDateTMP[1])

	openingRecordAtTMP := strings.Split(menu.OpeningRecordAt, ".")
	openingRecordAtTMPSecs, _ := strconv.Atoi(openingRecordAtTMP[0])
	openingRecordAtTMPNSecs, _ := strconv.Atoi(openingRecordAtTMP[1])

	closingRecordAtTMP := strings.Split(menu.ClosingRecordAt, ".")
	closingRecordAtTMPSecs, _ := strconv.Atoi(closingRecordAtTMP[0])
	closingRecordAtTMPNSecs, _ := strconv.Atoi(closingRecordAtTMP[1])

	createdAtTMP := strings.Split(menu.CreatedAt, ".")
	createdAtTMPSecs, _ := strconv.Atoi(createdAtTMP[0])
	createdAtTMPNSecs, _ := strconv.Atoi(createdAtTMP[1])

	result := &pb.Menu{
		Uuid:            menu.Uuid,
		OnDate:          &timestamppb.Timestamp{Seconds: int64(onDateTMPSecs), Nanos: int32(onDateTMPNSecs)},
		OpeningRecordAt: &timestamppb.Timestamp{Seconds: int64(openingRecordAtTMPSecs), Nanos: int32(openingRecordAtTMPNSecs)},
		ClosingRecordAt: &timestamppb.Timestamp{Seconds: int64(closingRecordAtTMPSecs), Nanos: int32(closingRecordAtTMPNSecs)},
		Salads:          menu.Salads.ToGRPCModel(),
		Garnishes:       menu.Garnishes.ToGRPCModel(),
		Meats:           menu.Meats.ToGRPCModel(),
		Soups:           menu.Soups.ToGRPCModel(),
		Drinks:          menu.Drinks.ToGRPCModel(),
		Desserts:        menu.Desserts.ToGRPCModel(),
		CreatedAt:       &timestamppb.Timestamp{Seconds: int64(createdAtTMPSecs), Nanos: int32(createdAtTMPNSecs)},
	}

	return &pb.GetMenuResponse{Menu: result}, nil
}

func (s *RestaurantServer) CreateProduct(ctx context.Context, in *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	log.Println("PRODUCT: create product:\n", in)

	time := timestamppb.New(time.Now())
	product := models.Product{
		Name:        in.Name,
		Description: in.Description,
		Type:        pb.ProductType_value[in.Type.String()],
		Weight:      in.Weight,
		Price:       in.Price,
		CreatedAt:   fmt.Sprintf("%v.%v", time.Seconds, time.Nanos),
	}

	if err := store.DB.Model(&models.Product{}).Create(&product).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.CreateProductResponse{}, nil
}

func (s *RestaurantServer) GetProduct(ctx context.Context, in *pb.GetProductListRequest) (*pb.GetProductListResponse, error) {
	log.Println("PRODUCT: get products")

	var products []models.Product
	var result []*pb.Product

	if err := store.DB.Model(&models.Product{}).Find(&products).Error; err != nil {
		return nil, err
	}

	for _, item := range products {
		created_at_string := strings.Split(item.CreatedAt, ".")
		secs, _ := strconv.Atoi(created_at_string[0])
		nans, _ := strconv.Atoi(created_at_string[1])

		result = append(result, &pb.Product{
			Uuid:        item.Uuid,
			Name:        item.Name,
			Description: item.Description,
			Type:        pb.ProductType(item.Type),
			Weight:      item.Weight,
			Price:       item.Price,
			CreatedAt:   &timestamppb.Timestamp{Seconds: int64(secs), Nanos: int32(nans)},
		})
	}

	return &pb.GetProductListResponse{Result: result}, nil
}

func (s *RestaurantServer) GetUpToDateOrderList(ctx context.Context, in *pb.GetUpToDateOrderListRequest) (*pb.GetUpToDateOrderListResponse, error) {
	log.Println("ORDER: get order")
	return &pb.GetUpToDateOrderListResponse{}, nil
}
