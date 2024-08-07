package repository

import (
	"context"
	"github.com/rs/zerolog/log"
	"mnezerka/MySpots/server/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_SPOTS = "spots"

type spotsRepository struct {
	db *mongo.Database
}

func NewSpotsRepository(db *mongo.Database) entities.SpotsRepository {
	return &spotsRepository{
		db: db,
	}
}

func (sr *spotsRepository) Create(c context.Context, spot *entities.Spot) error {
	log.Info().Str("module", "SpotsRepository").Msgf("creating new spot %v", spot)

	_, err := sr.db.Collection(COLLECTION_SPOTS).InsertOne(c, spot)
	return err
}

func (sr *spotsRepository) Fetch(c context.Context) ([]entities.Spot, error) {

	var spots []entities.Spot

	cursor, err := sr.db.Collection(COLLECTION_SPOTS).Find(c, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(c, &spots)
	if spots == nil {
		return []entities.Spot{}, err
	}

	return spots, err
}
