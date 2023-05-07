package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/models"
	pb "github.com/MSFT/pkg/services/restaurant"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertProductResponse(items []*pb.Product) []any {
	var products []any
	for _, product := range items {
		products = append(products, map[string]any{
			"uuid":        product.Uuid,
			"name":        product.Name,
			"description": product.Description,
			"type":        pb.ProductType_name[int32(product.Type)],
			"weight":      product.Weight,
			"price":       product.Price,
			"created_at":  time.Unix(product.CreatedAt.Seconds, int64(product.CreatedAt.Nanos)).Format(time.RFC3339),
		})
	}
	return products
}

func MenuRequest(c *gin.Context) {
	config := cfg.GetConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", config.General_host, config.Restaurant_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    http.StatusServiceUnavailable,
			"message": "service is unavailable",
			"details": nil,
		})
		return
	}
	defer conn.Close()
	client := pb.NewMenuServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	switch c.Request.Method {
	case "GET":
		grpc_response, err := client.GetMenu(ctx, &pb.GetMenuRequest{OnDate: timestamppb.Now()})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    int(status.FromContextError(err).Code()),
				"message": err,
				"details": nil,
			})
			return
		}

		response := make(map[string]any)
		response["uuid"] = grpc_response.Menu.Uuid
		response["on_date"] = time.Unix(grpc_response.Menu.OnDate.Seconds, int64(grpc_response.Menu.OnDate.Nanos)).Format(time.RFC3339)
		response["opening_record_at"] = time.Unix(grpc_response.Menu.OpeningRecordAt.Seconds, int64(grpc_response.Menu.OpeningRecordAt.Nanos)).Format(time.RFC3339)
		response["closing_record_at"] = time.Unix(grpc_response.Menu.ClosingRecordAt.Seconds, int64(grpc_response.Menu.ClosingRecordAt.Nanos)).Format(time.RFC3339)
		response["created_at"] = time.Unix(grpc_response.Menu.CreatedAt.Seconds, int64(grpc_response.Menu.CreatedAt.Nanos)).Format(time.RFC3339)

		if grpc_response.Menu.Salads != nil {
			response["salads"] = convertProductResponse(grpc_response.Menu.Salads)
		}
		if grpc_response.Menu.Garnishes != nil {
			response["garnishes"] = convertProductResponse(grpc_response.Menu.Garnishes)
		}
		if grpc_response.Menu.Meats != nil {
			response["meats"] = convertProductResponse(grpc_response.Menu.Meats)
		}
		if grpc_response.Menu.Soups != nil {
			response["soups"] = convertProductResponse(grpc_response.Menu.Soups)
		}
		if grpc_response.Menu.Drinks != nil {
			response["drinks"] = convertProductResponse(grpc_response.Menu.Drinks)
		}
		if grpc_response.Menu.Desserts != nil {
			response["desserts"] = convertProductResponse(grpc_response.Menu.Desserts)
		}

		c.JSON(http.StatusOK, gin.H{
			"menu": response,
		})

	case "POST":
		var request models.CreateMenuRequest
		if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    int(status.FromContextError(err).Code()),
				"message": err,
				"details": nil,
			})
			return
		}

		r, err := client.CreateMenu(ctx, &pb.CreateMenuRequest{
			OnDate:          timestamppb.New(request.OnDate),
			OpeningRecordAt: timestamppb.New(request.OpeningRecordAt),
			ClosingRecordAt: timestamppb.New(request.ClosingRecordAt),
			Salads:          request.Salads,
			Garnishes:       request.Garnishes,
			Meats:           request.Meats,
			Soups:           request.Soups,
			Drinks:          request.Drinks,
			Desserts:        request.Desserts,
		})
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"code":    int(status.FromContextError(err).Code()),
				"message": err,
				"details": nil,
			})
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

func ProductRequest(c *gin.Context) {
	config := cfg.GetConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", config.General_host, config.Restaurant_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    int(status.FromContextError(err).Code()),
			"message": err,
			"details": nil,
		})
		return
	}
	defer conn.Close()
	client := pb.NewProductServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	switch c.Request.Method {
	case "GET":
		r, err := client.GetProduct(ctx, &pb.GetProductListRequest{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    int(status.FromContextError(err).Code()),
				"message": err,
				"details": nil,
			})
			return
		}

		products := convertProductResponse(r.Result)

		c.JSON(http.StatusOK, gin.H{
			"result": products,
		})

	case "POST":
		var request map[string]any
		if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "request body is not parsed",
				"details": nil,
			})
			return
		}

		r, err := client.CreateProduct(ctx, &pb.CreateProductRequest{
			Name:        request["name"].(string),
			Description: request["description"].(string),
			Type:        pb.ProductType(pb.ProductType_value[request["type"].(string)]),
			Weight:      int32(request["weight"].(float64)),
			Price:       request["price"].(float64),
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    int(status.FromContextError(err).Code()),
				"message": err,
				"details": nil,
			})
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

func OrderRequest(c *gin.Context) {
	config := cfg.GetConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", config.General_host, config.Restaurant_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    int(status.FromContextError(err).Code()),
			"message": err,
			"details": nil,
		})
		return
	}
	defer conn.Close()
	client := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	r, err := client.GetUpToDateOrderList(ctx, &pb.GetUpToDateOrderListRequest{})
	if err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, r)
}
