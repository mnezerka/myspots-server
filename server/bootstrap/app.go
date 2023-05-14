package bootstrap

import (
	"log"
	"mnezerka/MySpots/server/mongo"
)

type Application struct {
	Env   *Env
	Mongo mongo.Client
}

func App() Application {

	log.Print("Creating application instance")

	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}
