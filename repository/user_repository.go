package repository

import (
	"context"
	"mnezerka/myspots-server/db"
	"mnezerka/myspots-server/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const COLLECTION_USERS = "users"

type UserRepository struct {
	db db.Database
}

func NewUserRepository(db db.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(c context.Context, user *entities.User) error {
	collection := ur.db.Collection(COLLECTION_USERS)

	_, err := collection.InsertOne(c, user)

	return err
}

func (ur *UserRepository) Fetch(c context.Context) ([]entities.User, error) {

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := ur.db.Collection(COLLECTION_USERS).Find(c, bson.D{}, opts)

	if err != nil {
		return nil, err
	}

	var users []entities.User

	err = cursor.All(c, &users)
	if users == nil {
		return []entities.User{}, err
	}

	return users, err
}

func (ur *UserRepository) GetByEmail(c context.Context, email string) (entities.User, error) {
	var user entities.User
	err := ur.db.Collection(COLLECTION_USERS).FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *UserRepository) GetByID(c context.Context, id string) (entities.User, error) {

	var user entities.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = ur.db.Collection(COLLECTION_USERS).FindOne(c, bson.M{"_id": idHex}).Decode(&user)
	return user, err
}
