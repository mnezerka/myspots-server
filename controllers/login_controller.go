package controllers

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"mnezerka/myspots-server/bootstrap"
	"mnezerka/myspots-server/entities"
	"net/http"
	"time"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginController struct {
	userRepository entities.UserRepository
	env            *bootstrap.Env
}

func NewLoginController(userRepository entities.UserRepository, env *bootstrap.Env) *LoginController {
	return &LoginController{userRepository: userRepository, env: env}
}

func (lc *LoginController) Login(c *gin.Context) {

	var request LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{Message: err.Error()})
		return
	}

	log.Debug().Str("module", "LoginController").Msgf("new login request for user '%s'", request.Email)

	// check if user exists
	user, err := lc.userRepository.GetByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, entities.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	log.Debug().Str("module", "LoginController").Msgf("user '%s' exists, validating password", request.Email)

	// user exists, check if password matches
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, entities.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	log.Debug().Str("module", "LoginController").Msgf("password for user '%s' is valid, generating token", request.Email)

	expirationTime := time.Now().Add(time.Hour * lc.env.AccessTokenExpiryHour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &entities.JwtCustomClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// generate new jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(lc.env.AccessTokenSecret))
	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: "error while encrypting token, try again"})
		return
	}

	loginResponse := LoginResponse{
		Token: tokenString,
	}

	c.JSON(http.StatusOK, loginResponse)
}
