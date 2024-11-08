package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jsanchezmen/microservices-proto/golang/payment"
	"github.com/jsanchezmen/microservices/payment/config"
	"github.com/jsanchezmen/microservices/payment/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	payment.UnimplementedPaymentServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", a.port))

	if err != nil {
		log.Fatalf("Failed to listen on port %v, error %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	payment.RegisterPaymentServer(grpcServer, a)

	if config.GetEnv() == "dev" {
		reflection.Register(grpcServer)
	}

	log.Printf("startomg payment service on port %d ...", a.port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve grpc on port %v", a.port)
	}

	log.Printf("GRPC Server started on port %v", a.port)
}
