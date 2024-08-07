package bootstrap

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func NewMongoDatabase(env *Env) *mongo.Database {

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

	log.Printf("Connecting to mongodb: %s", mongodbUri)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodbUri))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to mongodb: %s", mongodbUri)

	log.Printf("Trying to ping mongodb: %s", mongodbUri)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	/*
		defer func() {
			if err := client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
			log.Print("Connection to MongoDB closed.")
		}()
	*/

	return client.Database("myspots")
}
