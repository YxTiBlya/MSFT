package service

import (
	"context"
	"fmt"

	"github.com/MSFT/internal/cfg"
	log "github.com/MSFT/internal/log"
	restaurant_models "github.com/MSFT/internal/models/restaurant"
	"github.com/MSFT/pkg/services/customer"
	"github.com/MSFT/pkg/services/restaurant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UpdateOrderList(orders *restaurant_models.Orders, orderRequest *customer.CreateOrderRequest) error {
	config := cfg.GetConfig()

	conn_customer, err := grpc.Dial(fmt.Sprintf("%v:%d", config.Customer_host, config.Customer_grpc_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.ContextLogger.Error("dial customer service error:", err.Error())
		return err
	}
	defer conn_customer.Close()

	client_office := customer.NewOfficeServiceClient(conn_customer)
	response_office, err := client_office.GetOfficeByUUID(context.Background(), &customer.GetOfficeByUUIDRequest{OfficeUuid: orderRequest.OfficeUuid})
	if err != nil {
		log.ContextLogger.Error("UpdateOrderList error to get office by uuid:", err.Error())
		return err
	}
	office := response_office.Result

	conn_restaurant, err := grpc.Dial(fmt.Sprintf("%v:%d", config.Restaurant_host, config.Restaurant_grpc_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.ContextLogger.Error("dial restaurant service error:", err.Error())
		return err
	}
	defer conn_restaurant.Close()

	client_product := restaurant.NewProductServiceClient(conn_restaurant)

	appendOrder(orders, office, orderRequest.Salads, client_product)
	appendOrder(orders, office, orderRequest.Garnishes, client_product)
	appendOrder(orders, office, orderRequest.Meats, client_product)
	appendOrder(orders, office, orderRequest.Soups, client_product)
	appendOrder(orders, office, orderRequest.Drinks, client_product)
	appendOrder(orders, office, orderRequest.Desserts, client_product)

	return nil
}

func appendOrder(orders *restaurant_models.Orders, orderOffice *customer.Office, order []*customer.OrderItem, client_product restaurant.ProductServiceClient) {
	for _, orderItem := range order {
		response_product, err := client_product.GetProductByUUID(context.Background(), &restaurant.GetProductByUUIDRequest{ProductUuid: orderItem.ProductUuid})
		if err != nil {
			log.ContextLogger.Error("UpdateOrderList error to get product by uuid:", err)
			continue
		}
		product := response_product.Result

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
