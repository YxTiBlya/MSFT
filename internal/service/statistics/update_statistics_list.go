package service

import (
	"context"
	"fmt"

	"github.com/MSFT/internal/cfg"
	log "github.com/MSFT/internal/log"
	statistics_models "github.com/MSFT/internal/models/statistics"
	"github.com/MSFT/pkg/services/customer"
	"github.com/MSFT/pkg/services/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UpdateStatisticsList(statistics *statistics_models.Statistics, orderRequest *customer.CreateOrderRequest) error {
	config := cfg.GetConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", config.Restaurant_host, config.Restaurant_grpc_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.ContextLogger.Error("UpdateStatisticsList error:", err.Error())
		return err
	}
	defer conn.Close()

	client := restaurant.NewProductServiceClient(conn)

	updateStatistics(statistics, orderRequest.Salads, client)
	updateStatistics(statistics, orderRequest.Garnishes, client)
	updateStatistics(statistics, orderRequest.Meats, client)
	updateStatistics(statistics, orderRequest.Soups, client)
	updateStatistics(statistics, orderRequest.Drinks, client)
	updateStatistics(statistics, orderRequest.Desserts, client)

	return nil
}

func updateStatistics(statistics *statistics_models.Statistics, order []*customer.OrderItem, client restaurant.ProductServiceClient) {
	for _, orderItem := range order {
		response, err := client.GetProductByUUID(context.Background(), &restaurant.GetProductByUUIDRequest{ProductUuid: orderItem.ProductUuid})
		if err != nil {
			log.ContextLogger.Error("GetProductByUUID error:", err.Error())
			continue
		}
		product := response.Result

		statistics.Profit += float64(orderItem.Count) * product.Price

		productIsFind := false
		for idx, item := range statistics.TopProducts {
			if item.Uuid == product.Uuid {
				statistics.TopProducts[idx].Count += int(orderItem.Count)
				productIsFind = true
				break
			}
		}
		if !productIsFind {
			statistics.TopProducts = append(statistics.TopProducts, statistics_models.Product{
				Uuid:  product.Uuid,
				Name:  product.Name,
				Count: int(orderItem.Count),
				Type:  int32(product.Type),
			})
		}
	}
}
