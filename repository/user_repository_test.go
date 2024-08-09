package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mnezerka/myspots-server/entities"
	"mnezerka/myspots-server/mocks/db"
	"mnezerka/myspots-server/repository"
)

func TestCreate(t *testing.T) {

	mockDb := &db.MockDatabase{}
	mockCollection := &db.MockCollection{}

	//collectionName := entities.CollectionUser

	mockUser := &entities.User{
		ID:       primitive.NewObjectID(),
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "password",
	}

	emptyUser := &entities.User{}
	mockUserID := primitive.NewObjectID()

	t.Run("create user", func(t *testing.T) {

		mockDb.On("Collection", "users").Return(mockCollection)

		mockCollection.On("InsertOne", mock.Anything, mock.AnythingOfType("*entities.User")).Return(mockUserID, nil).Once()

		userRepository := repository.NewUserRepository(mockDb)

		err := userRepository.Create(context.Background(), mockUser)

		assert.NoError(t, err)

		mockCollection.AssertExpectations(t)
	})

	t.Run("fail for creation of empty user", func(t *testing.T) {
		mockDb.On("Collection", "users").Return(mockCollection)

		mockCollection.On("InsertOne", mock.Anything, mock.AnythingOfType("*entities.User")).Return(emptyUser, errors.New("Unexpected")).Once()

		ur := repository.NewUserRepository(mockDb)

		err := ur.Create(context.Background(), emptyUser)

		assert.Error(t, err)

		mockCollection.AssertExpectations(t)
	})
}
