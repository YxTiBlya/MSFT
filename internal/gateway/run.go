package gateway

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/internal/rabbitmq"
	"github.com/MSFT/internal/store"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server interface {
	RunGRPCServer(*cfg.Config, *grpc.Server)
	RuntHTTPServer(context.Context, *cfg.Config, *runtime.ServeMux)
}

func Run(cfg *cfg.Config, server server, recieve_broker bool) error {
	// create the connection to db
	if _, err := store.ConnToDB(cfg); err != nil {
		panic("failed to connect database:\n" + err.Error())
	}

	// create the connection to rabbitmq
	if err := rabbitmq.ConnToRabbitMQ(cfg); err != nil {
		panic("failed to connect rabbitmq:\n" + err.Error())
	}

	// logger init
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	if cfg.Logging_in_file {
		logger_file, err := os.OpenFile("logger/restaurant.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic("failed to create or opening the logger file:\n" + err.Error())
		}
		defer logger_file.Close()
		log.SetOutput(logger_file)
	}

	// server init
	s := grpc.NewServer()
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	go server.RunGRPCServer(cfg, s)
	go server.RuntHTTPServer(ctx, cfg, mux)
	if recieve_broker {
		go rabbitmq.RecieveOrder(cfg)
	}

	gracefulShutDown(s, cancel)

	return nil
}

func gracefulShutDown(s *grpc.Server, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done")
	log.Errorln(errorMessage)
	s.GracefulStop()
	cancel()
}
