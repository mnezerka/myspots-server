package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// spot represents our REST resource
type Spot struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name"`
	Created int32              `json:"created"`
}

type Spots struct {
	db *mongo.Database
}

func NewSpots(db *mongo.Database) *Spots {
	spots := &Spots{}
	spots.db = db

	return spots
}

func (s *Spots) CreateSpot(name string) (*Spot, error) {

	log.Printf("Creating spot '%s'", name)

	spot := &Spot{
		Name:    name,
		Created: int32(time.Now().Unix()),
	}

	collection := s.db.Collection("spots")

	result, err := collection.InsertOne(context.TODO(), spot)
	if err != nil {
		return nil, fmt.Errorf("error while creating spot %v", err)
	}

	spot.Id = result.InsertedID.(primitive.ObjectID)

	log.Printf("Created spot: %s (%s)", spot.Name, spot.Id.Hex())

	return spot, nil
}

func (s *Spots) List() ([]Spot, error) {

	collection := s.db.Collection(dbCollection)

	ctx := context.TODO()

	var result []Spot

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("db error (collection find): %v", err)
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		// To decode into a struct, use cursor.Decode()
		spot := Spot{}
		err := cur.Decode(&spot)
		if err != nil {
			log.Printf("db error (cursor decode): %v", err)
			return nil, err
		}
		result = append(result, spot)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
