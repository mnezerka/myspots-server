package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"mnezerka/myspots-server/bootstrap"
	"mnezerka/myspots-server/controllers"
	"mnezerka/myspots-server/middleware"
	"mnezerka/myspots-server/repository"
)

func main() {

	env := bootstrap.NewEnv()

	db := bootstrap.NewMongoDatabase(env)

	userRepository := repository.NewUserRepository(db)
	spotsRepository := repository.NewSpotsRepository(db)

	loginController := controllers.NewLoginController(userRepository, env)
	signupController := controllers.NewSignupController(userRepository, env)
	spotsController := controllers.NewSpotsController(spotsRepository)

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// r := gin.Default()

	//debugPrintWARNINGDefault()
	r := gin.New()
	r.Use(middleware.DefaultStructuredLogger()) // adds our structured logger
	r.Use(gin.Recovery())                       // adds the default recovery middleware

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.POST("/signup", signupController.Signup)
	r.POST("/login", loginController.Login)

	//r.GET("/spots", middleware.RequireAuth, spotsController.Fetch)
	r.Use(middleware.Authenticate(env))
	r.GET("/spots", spotsController.Fetch)
	r.POST("/spots", spotsController.Create)

	r.Run() // listen and serve on 0.0.0.0:8080
}
