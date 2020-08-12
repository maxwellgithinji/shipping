package main

import (
	"context"
	"fmt"
	"log"
	"os"

	micro "github.com/micro/go-micro/v2"

	pb "github.com/maxwellgithinji/shipping/shipping-service-vessel/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	service := micro.NewService(
		micro.Name("shipping.service.vessel"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shipping").Collection("vessels")

	repository := &MongoRepository{vesselCollection}

	h := &handler{repository}

	// Register handlers
	pb.RegisterVesselServiceHandler(service.Server(), h)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
