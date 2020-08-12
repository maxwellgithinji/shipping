package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vesselProto "github.com/maxwellgithinji/shipping/shipping-service-vessel/proto/vessel"
	micro "github.com/micro/go-micro/v2"

	pb "github.com/maxwellgithinji/shipping/shipping-service-consignment/proto/consignment"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	// Create a new service. Optionally include some options here.
	service := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shipping.service.consignment"),
	)

	// Init will parse the command line flags.
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

	consignmentCollection := client.Database("shipping").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}

	vesselClient := vesselProto.NewVesselService("shipping.service.client", service.Client())
	h := &handler{repository, vesselClient}

	// Register handlers
	pb.RegisterShippingServiceHandler(service.Server(), h)

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
