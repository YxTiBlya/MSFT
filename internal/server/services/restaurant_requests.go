package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/MSFT/internal/cfg"
	pb "github.com/MSFT/pkg/services/restaurant"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetMenuRequest(c *gin.Context) {
	config := cfg.GetConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", config.General_host, config.Restaurant_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}
	defer conn.Close()
	client := pb.NewMenuServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	r, err := client.GetMenu(ctx, &pb.GetMenuRequest{OnDate: timestamppb.Now()})
	if err != nil {
		return
	}
	fmt.Println(r.GetMenu())
}
