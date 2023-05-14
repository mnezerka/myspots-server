package controller

import (
	"mnezerka/MySpots/server/domain"
	"mnezerka/MySpots/server/internal/spatialutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpotsController struct {
	SpotsUsecase domain.SpotsUsecase
}

func (sc *SpotsController) Create(c *gin.Context) {
	var spot domain.Spot

	err := c.ShouldBind(&spot)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = spatialutil.ValidateCoordinates(spot.Coordinates)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	spot.ID = primitive.NewObjectID()

	spot.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	spot.Type = "Point"

	err = sc.SpotsUsecase.Create(c, &spot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Spot created successfully",
	})
}

func (u *SpotsController) Fetch(c *gin.Context) {
	//userID := c.GetString("x-user-id")

	tasks, err := u.SpotsUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
