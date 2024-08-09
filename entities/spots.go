package entities

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Spot struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" form:"name" binding:"required" json:"name"`
	Description string             `bson:"description" form:"description" json:"description"`
	UserID      primitive.ObjectID `bson:"userID" binding:"required" json:"-"`
	Coordinates Coordinates        `bson:"coordinates" form:"coordinates" binding:"required" json:"coordinates"`
}

type SpotCreateRequest struct {
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description"`
	Coordinates Coordinates `json:"coordinates" binding:"required"`
}

type SpotsRepository interface {
	Create(c context.Context, task *Spot) error
	Fetch(c context.Context) ([]Spot, error)
}
