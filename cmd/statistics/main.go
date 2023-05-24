package main

import (
	"github.com/BurntSushi/toml"
	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/gateway"
	gateway_statistics "github.com/MSFT/internal/gateway/statistics"
)

func main() {
	// parse and update config
	config := cfg.GetConfig()
	if _, err := toml.DecodeFile("config.toml", config); err != nil {
		panic("failed to decode toml file:\n" + err.Error())
	}
	config.Current_service = "statistics"
	cfg.UpdateConfig(config)

	// serve
	if err := gateway.Run(config, &gateway_statistics.StatisticsServer{}, true); err != nil {
		panic("error running gateway server:\n" + err.Error())
	}
}
