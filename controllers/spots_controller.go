package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mnezerka/myspots-server/entities"
	"mnezerka/myspots-server/internal/spatialutil"
	"net/http"
)

type SpotsController struct {
	spotsRepository entities.SpotsRepository
}

func NewSpotsController(spotsRepository entities.SpotsRepository) *SpotsController {
	return &SpotsController{spotsRepository: spotsRepository}
}

func (sc *SpotsController) Create(c *gin.Context) {
	var spotRequest entities.SpotCreateRequest

	log.Debug().Str("module", "SpotsController").Msg("create new spot")

	err := c.ShouldBind(&spotRequest)
	if err != nil {
		log.Warn().Str("module", "SpotsController").Msg("failed to parse request body")
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{Message: err.Error()})
		return
	}

	log.Debug().Str("module", "SpotsController").Msgf("parsed create request: %v", spotRequest)

	err = spatialutil.ValidateCoordinates(spotRequest.Coordinates)
	if err != nil {
		log.Warn().Str("module", "SpotsController").Msg("failed to validate coordinates")
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{Message: err.Error()})
		return
	}

	log.Debug().Str("module", "SpotsController").Msg("coordinates validated")

	userId := c.GetString("user-id")

	log.Debug().Str("module", "SpotsController").Msgf("current user id: %s", userId)

	var spot entities.Spot

	spot.ID = primitive.NewObjectID()
	spot.UserID, err = primitive.ObjectIDFromHex(userId)
	if err != nil {
		log.Warn().Str("module", "SpotsController").Msg("failed to convert user id to hex")
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{Message: err.Error()})
		return
	}

	spot.Name = spotRequest.Name
	spot.Description = spotRequest.Description
	spot.Coordinates = spotRequest.Coordinates
	log.Debug().Msgf("creating new spot: %v", spot)

	err = sc.spotsRepository.Create(c, &spot)
	if err != nil {
		log.Warn().Str("module", "SpotsController").Msgf("failed to create spot: %s", err.Error())
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	log.Debug().Msg("spot created successfully")

	c.JSON(http.StatusOK, entities.SuccessResponse{
		Message: "Spot created successfully",
	})
}

func (u *SpotsController) Fetch(c *gin.Context) {

	log.Debug().Str("module", "SpotsController").Msg("get spots")

	spots, err := u.spotsRepository.Fetch(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, spots)
}
