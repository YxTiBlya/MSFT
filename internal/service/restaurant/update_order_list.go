package service

import (
	customer_models "github.com/MSFT/internal/models/customer"
	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/internal/store"
	"github.com/MSFT/pkg/services/customer"
	log "github.com/sirupsen/logrus"
)

func UpdateOrderList(orders *restaurant_models.Orders, orderRequest *customer.CreateOrderRequest) error {
	var user customer_models.User
	if err := store.DB.Model(&customer_models.User{}).Where("uuid = ?", orderRequest.UserUuid).First(&user).Error; err != nil {
		return err
	}

	var office customer_models.Office
	if err := store.DB.Model(&customer_models.Office{}).Where("uuid = ?", user.Office_uuid).First(&office).Error; err != nil {
		return err
	}

	appendOrder(orders, office, orderRequest.Salads)
	appendOrder(orders, office, orderRequest.Garnishes)
	appendOrder(orders, office, orderRequest.Meats)
	appendOrder(orders, office, orderRequest.Soups)
	appendOrder(orders, office, orderRequest.Drinks)
	appendOrder(orders, office, orderRequest.Desserts)

	return nil
}

func appendOrder(orders *restaurant_models.Orders, orderOffice customer_models.Office, order []*customer.OrderItem) {
	var product restaurant_models.Product
	for _, orderItem := range order {
		if err := store.DB.Model(&restaurant_models.Product{}).Where("uuid = ?", orderItem.ProductUuid).First(&product).Error; err != nil {
			log.Errorln("product not finded:", err)
			continue
		}

		productIsFind := false
		for idx, totalItem := range orders.TotalOrders {
			if totalItem.ProductUuid == product.Uuid {
				orders.TotalOrders[idx].Count += int64(orderItem.Count)
				productIsFind = true
				break
			}
		}
		if !productIsFind {
			orders.TotalOrders = append(orders.TotalOrders, restaurant_models.Order{
				ProductUuid: product.Uuid,
				ProductName: product.Name,
				Count:       int64(orderItem.Count),
			})
		}

		orderOfficeIsFind := -1
		for idx, office := range orders.TotalOrdersByCompany {
			if office.CompanyUuid == orderOffice.Uuid {
				orderOfficeIsFind = idx
				break
			}
		}
		if orderOfficeIsFind != -1 {
			productIsFind = false
			for idx, totalItem := range orders.TotalOrdersByCompany[orderOfficeIsFind].Result {
				if totalItem.ProductUuid == product.Uuid {
					orders.TotalOrdersByCompany[orderOfficeIsFind].Result[idx].Count += int64(orderItem.Count)
					productIsFind = true
					break
				}
			}
			if !productIsFind {
				orders.TotalOrdersByCompany[orderOfficeIsFind].Result = append(orders.TotalOrdersByCompany[orderOfficeIsFind].Result, restaurant_models.Order{
					ProductUuid: product.Uuid,
					ProductName: product.Name,
					Count:       int64(orderItem.Count),
				})
			}
		} else {
			orders.TotalOrdersByCompany = append(orders.TotalOrdersByCompany, restaurant_models.OrderByOffice{
				CompanyUuid:   orderOffice.Uuid,
				OfficeName:    orderOffice.Name,
				OfficeAddress: orderOffice.Address,
				Result: restaurant_models.OrderArray{
					restaurant_models.Order{
						ProductUuid: product.Uuid,
						ProductName: product.Name,
						Count:       int64(orderItem.Count),
					},
				},
			})
		}
	}
}
