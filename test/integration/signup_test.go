package integration_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mnezerka/myspots-server/controllers"
	"mnezerka/myspots-server/entities"
)

func TestSignup(t *testing.T) {

	emptyUser := entities.User{}

	t.Run("with empty body", func(t *testing.T) {

		te := initTestEnv()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/signup", nil)
		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "missing form body")
	})

	t.Run("valid request for existing user", func(t *testing.T) {

		te := initTestEnv()

		mockUser := entities.User{
			ID:       primitive.NewObjectID(),
			Name:     "mn",
			Email:    "mn@gmail.com",
			Password: "mn",
		}

		te.mockUserRepository.On("GetByEmail", mock.AnythingOfType("*gin.Context"), "mn@example.com").Return(mockUser, nil).Once()

		data := controllers.SignupRequest{
			Name:     "mn",
			Email:    "mn@example.com",
			Password: "mn",
		}

		// marshall data to json (like json_encode)
		marshalled, err := json.Marshal(data)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(marshalled))
		req.Header.Set("Content-type", "application/json")
		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 409, w.Code)
		assert.Contains(t, w.Body.String(), "User already exists with the given email")
	})

	t.Run("valid request for new user", func(t *testing.T) {

		te := initTestEnv()

		data := controllers.SignupRequest{
			Name:     "mn",
			Email:    "mn@example.com",
			Password: "mn",
		}

		// mock request if user exists -> return error which means user doesn't exist
		te.mockUserRepository.On(
			"GetByEmail",
			mock.AnythingOfType("*gin.Context"),
			"mn@example.com").
			Return(emptyUser, errors.New("User not found")).
			Once()

		// mock request for creation of new user
		te.mockUserRepository.On(
			"Create",
			mock.AnythingOfType("*gin.Context"),
			mock.MatchedBy(func(u *entities.User) bool {
				return u.Name == "mn" && u.Email == "mn@example.com"
			})).
			Return(nil).
			Once()

		// marshall data to json (like json_encode)
		marshalled, err := json.Marshal(data)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/signup", bytes.NewReader(marshalled))
		req.Header.Set("Content-type", "application/json")
		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
	})

}
