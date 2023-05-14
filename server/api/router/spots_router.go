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

func NewSpotsRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	sr := repository.NewSpotsRepository(db, domain.CollectionSpots)
	sc := &controller.SpotsController{
		SpotsUsecase: usecase.NewSpotsUsecase(sr, timeout),
	}
	group.GET("/spots", sc.Fetch)
	group.POST("/spots", sc.Create)
}
