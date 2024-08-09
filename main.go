package main

import (
	"github.com/rs/zerolog"
	"mnezerka/myspots-server/bootstrap"
	"mnezerka/myspots-server/controllers"
	"mnezerka/myspots-server/repository"
	"mnezerka/myspots-server/router"
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

	r := router.SetupRouter(loginController, signupController, spotsController, env)

	r.Run() // listen and serve on 0.0.0.0:8080
}
