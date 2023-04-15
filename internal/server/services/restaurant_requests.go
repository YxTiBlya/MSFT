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

func MenuRequest(c *gin.Context) {
	config := cfg.GetConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", config.General_host, config.Restaurant_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
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
			return
		}

		c.JSON(http.StatusOK, r.GetMenu())

	case "POST":
		var request models.CreateMenuRequest
		if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
			c.JSON(http.StatusBadRequest, status.FromContextError(err))
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
			c.JSON(http.StatusBadRequest, status.FromContextError(err))
			return
		}

		c.JSON(http.StatusOK, r)
	}
}
