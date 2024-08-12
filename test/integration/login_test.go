package integration_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mnezerka/myspots-server/bootstrap"
	"mnezerka/myspots-server/controllers"
	"mnezerka/myspots-server/entities"
	mockentities "mnezerka/myspots-server/mocks/entities"
	"mnezerka/myspots-server/router"
)

func TestLogin(t *testing.T) {

	var r *gin.Engine
	var mockUserRepository *mockentities.MockUserRepository
	var env *bootstrap.Env
	emptyUser := entities.User{}

	var setup = func() {

		mockUserRepository = &mockentities.MockUserRepository{}

		env = &bootstrap.Env{}

		loginController := controllers.NewLoginController(mockUserRepository, env)

		r = router.SetupRouter(loginController, nil, nil, nil)
	}

	t.Run("with empty body", func(t *testing.T) {

		setup()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "missing form body")
	})

	t.Run("login of unkown user", func(t *testing.T) {

		setup()

		// mock request if user exists -> return error which means user doesn't exist
		mockUserRepository.On(
			"GetByEmail",
			mock.AnythingOfType("*gin.Context"),
			"mn@example.com").
			Return(emptyUser, errors.New("User not found")).
			Once()

		data := controllers.LoginRequest{
			Email:    "mn@example.com",
			Password: "pwd",
		}

		// marshall data to json (like json_encode)
		marshalled, err := json.Marshal(data)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", bytes.NewReader(marshalled))
		req.Header.Set("Content-type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
		assert.Contains(t, w.Body.String(), "User not found with the given email")
	})

	t.Run("valid login", func(t *testing.T) {

		setup()

		mockUser := entities.User{
			ID:       primitive.NewObjectID(),
			Name:     "mn",
			Email:    "mn@gmail.com",
			Password: "$2a$10$EqrtZNw/zjy/j2ZqI8Ne.u3rS3jgL/ufY3iCq0hLYcm/tIzWvTGqu",
		}

		// mock request if user exists -> return error which means user doesn't exist
		mockUserRepository.On(
			"GetByEmail",
			mock.AnythingOfType("*gin.Context"),
			"mn@example.com").
			Return(mockUser, nil).
			Once()

		data := controllers.LoginRequest{
			Email:    "mn@example.com",
			Password: "pwd",
		}

		// marshall data to json (like json_encode)
		marshalled, err := json.Marshal(data)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/login", bytes.NewReader(marshalled))
		req.Header.Set("Content-type", "application/json")
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		// check body
		var body map[string]interface{}
		assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &body))
		val, ok := body["token"]
		assert.True(t, ok)
		assert.NotEmpty(t, val)
	})

}
