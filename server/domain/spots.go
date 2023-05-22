package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionSpots = "spots"
)

type Spot struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" form:"name" binding:"required" json:"name"`
	Description string             `bson:"description" form:"description" binding:"required" json:"description"`
	UserID      primitive.ObjectID `bson:"userID" json:"-"`
	Type        string             `bson:"type" json:"-"`
	Coordinates Coordinates        `bson:"coordinates" form:"coordinates" binding:"required" json:"coordinates"`
}

type SpotsRepository interface {
	Create(c context.Context, task *Spot) error
	Fetch(c context.Context) ([]Spot, error)
}

type SpotsUsecase interface {
	Create(c context.Context, task *Spot) error
	Fetch(c context.Context) ([]Spot, error)
}
