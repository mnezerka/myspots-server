package router

import (
	"mnezerka/MySpots/server/api/controller"
	"mnezerka/MySpots/server/bootstrap"
	"mnezerka/MySpots/server/domain"
	"mnezerka/MySpots/server/mongo"
	"mnezerka/MySpots/server/repository"
	"mnezerka/MySpots/server/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
