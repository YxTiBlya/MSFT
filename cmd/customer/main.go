package main

import (
	"github.com/BurntSushi/toml"
	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/gateway"
	gateway_customer "github.com/MSFT/internal/gateway/customer"
)

func main() {
	// parse and update config
	config := cfg.GetConfig()
	if _, err := toml.DecodeFile("config.toml", config); err != nil {
		panic("failed to decode toml file:\n" + err.Error())
	}
	config.Current_service = "customer"
	cfg.UpdateConfig(config)

	// serve
	if err := gateway.Run(config, &gateway_customer.CustomerServer{}, false); err != nil {
		panic("error running gateway server:\n" + err.Error())
	}
}
