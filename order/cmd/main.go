package main

import (
	"log"

	"github.com/jsanchezmen/microservices/order/config"
	"github.com/jsanchezmen/microservices/order/internal/adapters/db"
	"github.com/jsanchezmen/microservices/order/internal/adapters/grpc"
	"github.com/jsanchezmen/microservices/order/internal/adapters/payment"
	"github.com/jsanchezmen/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())

	if err != nil {
		log.Fatalf("Fiailed to connecto database. error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())

	application := api.NewApplication(dbAdapter, paymentAdapter)
	defer paymentAdapter.Connection.Close()

	if err != nil {
		log.Fatalf("Fiailed to Init Payment Stub: %v", err)
	}

	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
