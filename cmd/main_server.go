package main

import (
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/MSFT/internal/cfg"
	customer_handlers "github.com/MSFT/internal/server/services/customer"
	restaurant_handlers "github.com/MSFT/internal/server/services/restaurant"
	"github.com/MSFT/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// parse and update config
	config := cfg.GetConfig()
	if _, err := toml.DecodeFile("config.toml", config); err != nil {
		panic("failed to decode toml file:\n" + err.Error())
	}
	cfg.UpdateConfig(config)

	// logger init
	logger_file, err := os.OpenFile("logger/main_server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("failed to create or opening the logger file:\n" + err.Error())
	}
	defer logger_file.Close()
	gin.DefaultWriter = io.MultiWriter(logger_file, os.Stdout)

	// urls
	r := gin.New()
	r.Use(middlewares.Logger(), gin.Recovery())

	// restaurant urls
	r.GET("/restaurant/menu", restaurant_handlers.MenuRequest)
	r.POST("/restaurant/menu", restaurant_handlers.MenuRequest)
	r.GET("/restaurant/orders", restaurant_handlers.OrderRequest)
	r.GET("/restaurant/products", restaurant_handlers.ProductRequest)
	r.POST("/restaurant/products", restaurant_handlers.ProductRequest)

	// customer urls
	r.GET("/customer/offices", customer_handlers.OfficeRequest)
	r.POST("/customer/offices", customer_handlers.OfficeRequest)

	// serve
	if err := r.Run(fmt.Sprintf(":%d", config.General_port)); err != nil {
		panic("failed to start server:\n" + err.Error())
	}
}
