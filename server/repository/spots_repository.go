package repository

import (
	"context"
	"mnezerka/MySpots/server/domain"
	"mnezerka/MySpots/server/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type spotsRepository struct {
	database   mongo.Database
	collection string
}

func NewSpotsRepository(db mongo.Database, collection string) domain.SpotsRepository {
	return &spotsRepository{
		database:   db,
		collection: collection,
	}
}

func (sr *spotsRepository) Create(c context.Context, spot *domain.Spot) error {
	collection := sr.database.Collection(sr.collection)

	_, err := collection.InsertOne(c, spot)

	return err
}

func (sr *spotsRepository) Fetch(c context.Context) ([]domain.Spot, error) {
	collection := sr.database.Collection(sr.collection)

	var spots []domain.Spot

	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &spots)
	if spots == nil {
		return []domain.Spot{}, err
	}

	return spots, err
}
