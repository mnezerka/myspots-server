package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mnezerka/myspots-server/entities"
)

func TestSpotsCreate(t *testing.T) {

	t.Run("check auth", func(t *testing.T) {

		te := initTestEnv()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/spots", nil)
		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 403, w.Code)
		assert.Contains(t, w.Body.String(), "missing authorization header")

		req, _ = http.NewRequest("GET", "/spots", nil)
		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 403, w.Code)
		assert.Contains(t, w.Body.String(), "missing authorization header")
	})

	t.Run("create with empty body", func(t *testing.T) {

		te := initTestEnv()

		token := te.login(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/spots", nil)
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "invalid request")
	})

	t.Run("create with wrong body", func(t *testing.T) {

		te := initTestEnv()

		token := te.login(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/spots", bytes.NewBuffer([]byte("this is wrong body")))
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "invalid character")
	})

	t.Run("create new spot", func(t *testing.T) {

		te := initTestEnv()

		// mock request to store new spot
		te.mockSpotsRepository.On(
			"Create",
			mock.AnythingOfType("*gin.Context"),
			mock.MatchedBy(func(s *entities.Spot) bool {
				return s.Name == "new-spot" &&
					s.Description == "new-spot-description" &&
					len(s.Coordinates) == 2 &&
					s.Coordinates[0] == 30.5 &&
					s.Coordinates[1] == 60.4
			})).
			Return(nil).
			Once()

		token := te.login(t)

		data := entities.SpotCreateRequest{
			Name:        "new-spot",
			Description: "new-spot-description",
			Coordinates: []float32{30.5, 60.4},
		}

		// marshall data to json (like json_encode)
		marshalled, err := json.Marshal(data)
		assert.Nil(t, err)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/spots", bytes.NewReader(marshalled))
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "Spot created successfully")
	})
}

func TestSpotsGet(t *testing.T) {

	t.Run("check auth", func(t *testing.T) {

		te := initTestEnv()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/spots", nil)
		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 403, w.Code)
		assert.Contains(t, w.Body.String(), "missing authorization header")
	})

	t.Run("get empty spots", func(t *testing.T) {

		te := initTestEnv()

		// mock request to store new spot
		te.mockSpotsRepository.On(
			"Fetch",
			mock.AnythingOfType("*gin.Context")).
			Return([]entities.Spot{}, nil).
			Once()

		token := te.login(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/spots", nil)
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		// check body
		var body []entities.Spot
		assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &body))
		assert.Empty(t, body)
	})

	t.Run("get spots", func(t *testing.T) {

		te := initTestEnv()

		spotId := primitive.NewObjectID()
		userId := primitive.NewObjectID()

		spot1 := entities.Spot{
			ID:          spotId,
			Name:        "spot1",
			Description: "spot1-desc",
			Coordinates: []float32{30.5, 60.4},
			UserID:      userId,
		}

		// mock request to store new spot
		te.mockSpotsRepository.On(
			"Fetch",
			mock.AnythingOfType("*gin.Context")).
			Return([]entities.Spot{spot1}, nil).
			Once()

		token := te.login(t)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/spots", nil)
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		te.ginEngine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		// check body
		var body []entities.Spot
		assert.Nil(t, json.Unmarshal(w.Body.Bytes(), &body))
		assert.Equal(t, 1, len(body))
		assert.Equal(t, "spot1", body[0].Name)
	})

}
