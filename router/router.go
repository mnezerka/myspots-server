package router

import (
	"github.com/gin-gonic/gin"
	"mnezerka/myspots-server/bootstrap"
	"mnezerka/myspots-server/controllers"
	"mnezerka/myspots-server/middleware"
)

func SetupRouter(loginController *controllers.LoginController,
	signupController *controllers.SignupController,
	spotsController *controllers.SpotsController,
	env *bootstrap.Env) *gin.Engine {
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

	return r
}
