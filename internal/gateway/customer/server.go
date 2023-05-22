package gateway_customer

import (
	"context"
	"fmt"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/MSFT/internal/cfg"
	service "github.com/MSFT/internal/service/customer"
	pb "github.com/MSFT/pkg/services/customer"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CustomerServer struct{}

func (rs *CustomerServer) RunGRPCServer(cfg *cfg.Config, s *grpc.Server) {
	service := &service.CustomerService{}
	pb.RegisterOfficeServiceServer(s, service)
	pb.RegisterOrderServiceServer(s, service)
	pb.RegisterUserServiceServer(s, service)

	l, err := net.Listen("tcp", fmt.Sprintf("%v:%d", cfg.General_host, cfg.Customer_grpc_service_port))
	if err != nil {
		log.Fatalln("failed to listen:\n" + err.Error())
	}

	log.Infof("starting listening grpc server at %v", fmt.Sprintf("%v:%d", cfg.General_host, cfg.Customer_grpc_service_port))
	if err := s.Serve(l); err != nil {
		log.Fatalln("error service grpc server:\n" + err.Error())
	}
}

func (rs *CustomerServer) RuntHTTPServer(ctx context.Context, cfg *cfg.Config, mux *runtime.ServeMux) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	endpoint := fmt.Sprintf("%v:%d", cfg.General_host, cfg.Customer_grpc_service_port)

	if err := pb.RegisterOfficeServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatalln(err)
	}

	if err := pb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatalln(err)
	}

	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatalln(err)
	}

	log.Infof("starting listening http server at %s", fmt.Sprintf("%v:%d", cfg.General_host, cfg.Customer_http_service_port))
	if err := http.ListenAndServe(fmt.Sprintf(":%v", cfg.Customer_http_service_port), mux); err != nil {
		log.Fatalln("error service http server:\n" + err.Error())
	}
}

func (rs *CustomerServer) RunRabbitMQReciever(cfg *cfg.Config) {

}
