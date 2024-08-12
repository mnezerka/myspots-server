package integration_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mnezerka/myspots-server/bootstrap"
	"mnezerka/myspots-server/controllers"
	"mnezerka/myspots-server/entities"
	mockentities "mnezerka/myspots-server/mocks/entities"
	"mnezerka/myspots-server/router"
)

func TestSpots(t *testing.T) {

	var r *gin.Engine
	var mockUserRepository *mockentities.MockUserRepository
	var mockSpotsRepository *mockentities.MockSpotsRepository
	var env *bootstrap.Env

	var setup = func() {

		mockUserRepository = &mockentities.MockUserRepository{}
		mockSpotsRepository = &mockentities.MockSpotsRepository{}

		env = &bootstrap.Env{
			TokenExpiryHour: 4 * time.Hour,
			TokenSecret:     "some-secret",
		}

		loginController := controllers.NewLoginController(mockUserRepository, env)
		spotsController := controllers.NewSpotsController(mockSpotsRepository)

		r = router.SetupRouter(loginController, nil, spotsController, env)
	}

	t.Run("check auth", func(t *testing.T) {

		setup()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/spots", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 403, w.Code)
		assert.Contains(t, w.Body.String(), "missing authorization header")

		req, _ = http.NewRequest("GET", "/spots", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 403, w.Code)
		assert.Contains(t, w.Body.String(), "missing authorization header")
	})

	t.Run("create with empty body", func(t *testing.T) {

		setup()

		token := login(t, r, mockUserRepository)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/spots", nil)
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "invalid request")
	})

	t.Run("create with wrong body", func(t *testing.T) {

		setup()

		token := login(t, r, mockUserRepository)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/spots", bytes.NewBuffer([]byte("this is wrong body")))
		req.Header.Set("Content-type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), "invalid character")
	})

	t.Run("create new spot", func(t *testing.T) {

		setup()

		// mock request to store new spot
		mockSpotsRepository.On(
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

		token := login(t, r, mockUserRepository)

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

		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "Spot created successfully")
	})
}
