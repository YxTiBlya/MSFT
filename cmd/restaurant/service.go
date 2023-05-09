package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/MSFT/internal/cfg"
	restaurant_handlers "github.com/MSFT/internal/server/services/restaurant"
	"github.com/MSFT/internal/store"
	pb "github.com/MSFT/pkg/services/restaurant"
	"google.golang.org/grpc"
)

func main() {
	// parse and update config
	config := cfg.GetConfig()
	if _, err := toml.DecodeFile("config.toml", config); err != nil {
		panic("failed to decode toml file:\n" + err.Error())
	}
	cfg.UpdateConfig(config)

	// create the connection to db
	if _, err := store.ConnToDB(config); err != nil {
		panic("failed to connect database:\n" + err.Error())
	}

	// logger init
	logger_file, err := os.OpenFile("logger/restaurant.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("failed to create or opening the logger file:\n" + err.Error())
	}
	defer logger_file.Close()
	//log.SetOutput(logger_file)

	// listener init
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%d", config.General_host, config.Restaurant_service_port))
	if err != nil {
		log.Fatalln("failed to listen:\n" + err.Error())
	}

	// init grpc server
	server := grpc.NewServer()
	server_model := restaurant_handlers.RestaurantServer{}
	pb.RegisterMenuServiceServer(server, &server_model)
	pb.RegisterOrderServiceServer(server, &server_model)
	pb.RegisterProductServiceServer(server, &server_model)

	// serve
	log.Printf("server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		panic("failed to start server:\n" + err.Error())
	}
}
