package main

import (
	"github.com/BurntSushi/toml"
	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/gateway"
	gateway_restaurant "github.com/MSFT/internal/gateway/restaurant"
)

func main() {
	// parse and update config
	config := cfg.GetConfig()
	if _, err := toml.DecodeFile("config.toml", config); err != nil {
		panic("failed to decode toml file:\n" + err.Error())
	}
	cfg.UpdateConfig(config)

	// serve
	if err := gateway.Run(config, &gateway_restaurant.RestaurantServer{}, true); err != nil {
		panic("error running gateway server:\n" + err.Error())
	}
}
