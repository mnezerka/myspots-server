package main

import (
	"log"
	"mnezerka/MySpots/server/api/router"
	"mnezerka/MySpots/server/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("Starting...")

	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	router.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}
