package service

import (
	restaurant_models "github.com/MSFT/internal/models/restaurant"
	statistics_models "github.com/MSFT/internal/models/statistics"
	"github.com/MSFT/internal/store"
	"github.com/MSFT/pkg/services/customer"
	log "github.com/sirupsen/logrus"
)

func UpdateStatisticsList(statistics *statistics_models.Statistics, orderRequest *customer.CreateOrderRequest) error {
	updateStatistics(statistics, orderRequest.Salads)
	updateStatistics(statistics, orderRequest.Garnishes)
	updateStatistics(statistics, orderRequest.Meats)
	updateStatistics(statistics, orderRequest.Soups)
	updateStatistics(statistics, orderRequest.Drinks)
	updateStatistics(statistics, orderRequest.Desserts)

	return nil
}

func updateStatistics(statistics *statistics_models.Statistics, order []*customer.OrderItem) {
	var product restaurant_models.Product
	for _, orderItem := range order {
		if err := store.DB.Model(&restaurant_models.Product{}).Where("uuid = ?", orderItem.ProductUuid).First(&product).Error; err != nil {
			log.Errorln("product not finded:", err)
			continue
		}
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
				Type:  product.Type,
			})
		}
	}
}
