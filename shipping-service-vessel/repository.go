package main

import (
	"context"

	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"

	pb "github.com/maxwellgithinji/shipping/shipping-service-vessel/proto/vessel"
)

type Vessel struct {
	ID        string `json:"id"`
	Capacity  int32  `json:"capacity"`
	MaxWeight int32  `json:"max_weight"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerID   string `json:"owner_id"`
}

type Specification struct {
	Capacity  int32 `json:"capacity"`
	MaxWeight int32 `json:"max_weight"`
}

type Response struct {
	Vessel  Vessel  `json:"vessel"`
	Vessels Vessels `json:"vessels"`
}

type Vessels []*Vessel

type repository interface {
	FindAvailable(ctx context.Context, spec *Specification) (error, *Vessel)
	Create(ctx context.Context, vessel *Vessel) error
}

//MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// UnmarshalSpecification is a helper func
func UnmarshalSpecification(spec *Specification) *pb.Specification {
	return &pb.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

// MarshalSpecification is a helper func
func MarshalSpecification(spec *pb.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

// UnmarshalVessel is a helper function
func UnmarshalVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerID,
	}
}

// MarshalVessel is a helper function
func MarshalVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerID:   vessel.OwnerId,
	}
}

// FindAvailable checks the specifications against a map of vessels
// if capacity and max weight are below a vessels capacity and max weight
// then returns that vessel
func (repository *MongoRepository) FindAvailable(ctx context.Context, spec *Specification) (error, *Vessel) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	vessel := &Vessel{}
	if err := repository.collection.FindOne(ctx, filter).Decode(vessel); err != nil {
		return err, nil
	}
	return nil, vessel
}

// Create a new vessel
func (repository *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err := repository.collection.InsertOne(ctx, vessel)
	return err
}
