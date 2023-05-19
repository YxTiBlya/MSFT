package gateway_restaurant

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/MSFT/internal/cfg"
	service "github.com/MSFT/internal/service/restaurant"
	pb "github.com/MSFT/pkg/services/restaurant"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RestaurantServer struct{}

func (rs *RestaurantServer) RunGRPCServer(cfg *cfg.Config, s *grpc.Server) {
	service := &service.RestaurantService{}
	pb.RegisterMenuServiceServer(s, service)
	pb.RegisterOrderServiceServer(s, service)
	pb.RegisterProductServiceServer(s, service)

	l, err := net.Listen("tcp", fmt.Sprintf("%v:%d", cfg.General_host, cfg.Restaurant_grpc_service_port))
	if err != nil {
		log.Fatalln("failed to listen:\n" + err.Error())
	}

	log.Printf("starting listening grpc server at %v", fmt.Sprintf("%v:%d", cfg.General_host, cfg.Restaurant_grpc_service_port))
	if err := s.Serve(l); err != nil {
		panic("error service grpc server:\n" + err.Error())
	}
}

func (rs *RestaurantServer) RuntHTTPServer(ctx context.Context, cfg *cfg.Config, mux *runtime.ServeMux) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := fmt.Sprintf("%v:%d", cfg.General_host, cfg.Restaurant_grpc_service_port)

	if err := pb.RegisterMenuServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatal(err)
	}

	if err := pb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatal(err)
	}

	if err := pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatal(err)
	}

	log.Printf("starting listening http server at %s", fmt.Sprintf("%v:%d", cfg.General_host, cfg.Restaurant_http_service_port))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Restaurant_http_service_port), mux); err != nil {
		log.Fatalf("error service http server %v", err)
	}
}
