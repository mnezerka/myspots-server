package router

import (
	"time"

	"mnezerka/MySpots/server/api/middleware"
	"mnezerka/MySpots/server/bootstrap"
	"mnezerka/MySpots/server/mongo"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {

	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	//NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	//	NewProfileRouter(env, timeout, db, protectedRouter)
	NewSpotsRouter(env, timeout, db, protectedRouter)

	gin.StaticFile("/", "./webjs/index.html")
	gin.StaticFile("/spots.js", "./webjs/spots.js")
}
