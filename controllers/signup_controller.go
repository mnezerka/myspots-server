package controllers

import (
	"mnezerka/myspots-server/bootstrap"
	"mnezerka/myspots-server/entities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignupController struct {
	userRepository entities.UserRepository
	env            *bootstrap.Env
}

func NewSignupController(userRepository entities.UserRepository, env *bootstrap.Env) *SignupController {
	return &SignupController{userRepository: userRepository, env: env}
}

func (sc *SignupController) Signup(c *gin.Context) {

	var request SignupRequest

	//body, _ := io.ReadAll(c.Request.Body)
	//log.Debug().Str("module", "SignupController").Msgf("Incoming http request: %s", body)

	log.Debug().Str("module", "SignupController").Msg("Incoming http request")

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = sc.userRepository.GetByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, entities.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	log.Debug().Str("module", "SignupController").Msgf("user with email %s does not exist -> could be created", request.Email)

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	log.Debug().Str("module", "SignupController").Msgf("user password successfully encrypted")

	request.Password = string(encryptedPassword)

	user := entities.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	log.Debug().Str("module", "SignupController").Msgf("creating new user instance")

	err = sc.userRepository.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}
