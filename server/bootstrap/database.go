package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"mnezerka/MySpots/server/mongo"
)

func NewMongoDatabase(env *Env) mongo.Client {

	log.Print("Opening and initializing mongo database")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	var mongodbUri string

	if dbUser == "" || dbPass == "" {
		mongodbUri = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	} else {
		mongodbUri = fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)
	}

	client, err := mongo.NewClient(mongodbUri)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connection to MongoDB closed.")
}
