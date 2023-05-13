package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUri = "mongodb://localhost:27017"
const dbName = "spots"
const dbCollection = "spots"

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func main() {
	// try to open database
	log.Printf("Opening database '%s'", dbUri)
	dbClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUri))
	if err != nil {
		log.Fatalf("Failed to open database on %s (%v)", dbUri, err)
	}

	// check the connection
	log.Print("Checking database connection")
	err = dbClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Cannot ping database on %s (%v)", dbUri, err)
		os.Exit(1)
	}

	// look for collection
	log.Printf("Setting default database to '%s'", dbName)
	db := dbClient.Database(dbName)

	// creating collection if it doesn't exist
	log.Printf("Looking for collections: '%s'", dbCollection)
	cNames, err := db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		log.Fatalf("db error (list collections): %s", err)
	}

	cSpotsExists := false
	for _, cName := range cNames {
		if cName == dbCollection {
			cSpotsExists = true
		}
	}

	if !cSpotsExists {
		log.Printf("Creating collection '%s'", dbCollection)
		db.CreateCollection(context.TODO(), dbCollection)
	}

	spotH := &spotHandler{
		db:    db,
		spots: NewSpots(db),
	}

	mux := http.NewServeMux()
	mux.Handle("/api/v1/spots", spotH)
	mux.Handle("/api/v1/spots/", spotH)

	log.Print("Listening...")
	http.ListenAndServe("localhost:8080", mux)
}
