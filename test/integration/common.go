package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mnezerka/myspots-server/controllers"
	"mnezerka/myspots-server/entities"
	mockentities "mnezerka/myspots-server/mocks/entities"
)

func login(t *testing.T, r *gin.Engine, mockUserRepository *mockentities.MockUserRepository) string {

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

	// login to get token
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

	return body["token"].(string)
}
