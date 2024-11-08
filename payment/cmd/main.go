package main

import (
	"log"

	"github.com/jsanchezmen/microservices/payment/config"
	"github.com/jsanchezmen/microservices/payment/internal/adapters/db"
	"github.com/jsanchezmen/microservices/payment/internal/adapters/grpc"
	"github.com/jsanchezmen/microservices/payment/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatal(err)
	}
	application := api.NewApplication(dbAdapter)

	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())

	grpcAdapter.Run()
}
