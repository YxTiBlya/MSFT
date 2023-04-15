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
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
		r, err := client.GetMenu(ctx, &pb.GetMenuRequest{OnDate: timestamppb.Now()})
		if err != nil {
			//c.JSON(http.StatusBadRequest, status.FromContextError(err))
			return
		}

		c.JSON(http.StatusOK, r)

	case "POST":
		var request models.CreateMenuRequest
		if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "request body is not parsed",
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
			//c.JSON(http.StatusBadRequest, status.FromContextError(err))
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
			"code":    http.StatusServiceUnavailable,
			"message": "service is unavailable",
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
			//c.JSON(http.StatusBadRequest, gin.H{})
			return
		}

		c.JSON(http.StatusOK, r)

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
			//
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
			"code":    http.StatusServiceUnavailable,
			"message": "service is unavailable",
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
