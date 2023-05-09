package customer_handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/MSFT/internal/cfg"
	pb "github.com/MSFT/pkg/services/customer"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func OfficeRequest(c *gin.Context) {
	config := cfg.GetConfig()

	conn, err := grpc.Dial(fmt.Sprintf("%v:%d", config.General_host, config.Customer_service_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    http.StatusServiceUnavailable,
			"message": "service is unavailable",
			"details": nil,
		})
		return
	}
	defer conn.Close()
	client := pb.NewOfficeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	switch c.Request.Method {
	case "GET":
		grpc_response, err := client.GetOfficeList(ctx, &pb.GetOfficeListRequest{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    int(status.FromContextError(err).Code()),
				"message": err,
				"details": nil,
			})
			return
		}

		c.JSON(http.StatusOK, grpc_response)
	case "POST":
		grpc_response, err := client.CreateOffice(ctx, &pb.CreateOfficeRequest{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    int(status.FromContextError(err).Code()),
				"message": err,
				"details": nil,
			})
			return
		}

		c.JSON(http.StatusOK, grpc_response)
	}
}

func OrderRequest(c *gin.Context) {

}

func UserRequest(c *gin.Context) {

}
