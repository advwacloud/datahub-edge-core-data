package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	//"core-data/clients"
	"core-data/handler"
	"core-data/subscriber"

	core "github.com/advwacloud/datahub-edge-domain-models/protos/core-data"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("datahub.edge.core-data"),
		micro.Version("beta"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	core.RegisterCoreDataHandler(service.Server(), new(handler.CoreData))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("datahub.edge.core-data", service.Server(), new(subscriber.CoreData))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
