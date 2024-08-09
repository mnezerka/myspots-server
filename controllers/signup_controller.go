package controllers

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"mnezerka/myspots-server/bootstrap"
	"mnezerka/myspots-server/entities"
	"mnezerka/myspots-server/repository"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *entities.User) error
	GetUserByEmail(c context.Context, email string) (entities.User, error)
	CreateAccessToken(user *entities.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *entities.User, secret string, expiry int) (refreshToken string, err error)
}

type SignupController struct {
	userRepository *repository.UserRepository
	env            *bootstrap.Env
}

func NewSignupController(userRepository *repository.UserRepository, env *bootstrap.Env) *SignupController {
	return &SignupController{userRepository: userRepository, env: env}
}

func (sc *SignupController) Signup(c *gin.Context) {

	var request SignupRequest

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

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := entities.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.userRepository.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entities.ErrorResponse{Message: err.Error()})
		return
	}

	// Generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	accessToken, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, entities.ErrorResponse{Message: "Failed to create token"})
		return
	}

	signupResponse := SignupResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
