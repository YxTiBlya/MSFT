package gateway

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MSFT/internal/cfg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server interface {
	RunGRPCServer(*cfg.Config, *grpc.Server)
	RuntHTTPServer(context.Context, *cfg.Config, *runtime.ServeMux)
}

func Run(cfg *cfg.Config, server server) error {
	s := grpc.NewServer()
	mux := runtime.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())

	go server.RunGRPCServer(cfg, s)
	go server.RuntHTTPServer(ctx, cfg, mux)

	gracefulShutDown(s, cancel)

	return nil
}

func gracefulShutDown(s *grpc.Server, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done")
	log.Println(errorMessage)
	s.GracefulStop()
	cancel()
}
