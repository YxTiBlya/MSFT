package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/gateway"
	gateway_restaurant "github.com/MSFT/internal/gateway/restaurant"
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

	// logger init
	logger_file, err := os.OpenFile("logger/restaurant.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("failed to create or opening the logger file:\n" + err.Error())
	}
	defer logger_file.Close()
	//log.SetOutput(logger_file)

	// serve
	if err := gateway.Run(config, &gateway_restaurant.RestaurantServer{}); err != nil {
		log.Fatal("error running gateway server ", err)
	}
}
