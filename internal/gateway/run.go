package gateway

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MSFT/internal/cfg"
	log "github.com/MSFT/internal/log"
	"github.com/MSFT/internal/rabbitmq"
	"github.com/MSFT/internal/store"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server interface {
	RunGRPCServer(*cfg.Config, *grpc.Server)
	RuntHTTPServer(context.Context, *cfg.Config, *runtime.ServeMux)
	RunRabbitMQReciever(c *cfg.Config)
}

func Run(cfg *cfg.Config, server server, recieve_broker bool) error {
	// logger init
	log.InitLogger(cfg)

	// create the connection to db
	if _, err := store.ConnToDB(cfg); err != nil {
		log.ContextLogger.Fatal("failed to connect database:", err.Error())
	}

	// create the connection to rabbitmq
	if err := rabbitmq.ConnToRabbitMQ(cfg); err != nil {
		log.ContextLogger.Fatal("failed to connect rabbitmq:", err.Error())
	}

	// server init
	s := grpc.NewServer()
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	go server.RunGRPCServer(cfg, s)
	go server.RuntHTTPServer(ctx, cfg, mux)
	if recieve_broker {
		go server.RunRabbitMQReciever(cfg)
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
	log.ContextLogger.Error(errorMessage)
	s.GracefulStop()
	cancel()
}
