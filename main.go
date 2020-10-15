package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	//"core-data/clients"
	"core-data/handler"
	"core-data/subscriber"

	core "github.com/advwacloud/datahub-edge-domain-models/protos/core-data"
	//coredata "core-data/proto/core-data"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.core-data"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	core.RegisterCoreDataHandler(service.Server(), new(handler.CoreData))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.core-data", service.Server(), new(subscriber.CoreData))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
