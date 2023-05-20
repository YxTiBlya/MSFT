package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/gateway"
	gateway_customer "github.com/MSFT/internal/gateway/customer"
	"github.com/MSFT/internal/rabbitmq"
	"github.com/MSFT/internal/store"
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

	// create the connection to rabbitmq
	if err := rabbitmq.ConnToRabbitMQ(config); err != nil {
		panic("failed to connect rabbitmq:\n" + err.Error())
	}

	// logger init
	if config.Logging_in_file {
		logger_file, err := os.OpenFile("logger/customer.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("failed to create or opening the logger file:\n" + err.Error())
		}
		defer logger_file.Close()
		log.SetOutput(logger_file)
	}

	// serve
	if err := gateway.Run(config, &gateway_customer.CustomerServer{}, false); err != nil {
		log.Fatal("error running gateway server ", err)
	}
}
