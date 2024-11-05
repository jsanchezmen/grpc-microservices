package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jsanchezmen/microservices-proto/golang/order"
	"github.com/jsanchezmen/microservices/order/config"
	"github.com/jsanchezmen/microservices/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.ApiPort
	port   int
	server *grpc.Server
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.ApiPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	a.server = grpcServer
	order.RegisterOrderServer(grpcServer, a)

	if config.GetEnv() == "dev" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve grpc on port %v", a.port)
	}

}
