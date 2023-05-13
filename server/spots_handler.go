package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	listSpotsRe  = regexp.MustCompile(`^\/api/v1/spots[\/]*$`)
	getSpotRe    = regexp.MustCompile(`^\/api/v1/spots\/(\d+)$`)
	createSpotRe = regexp.MustCompile(`^\/api/v1/spots[\/]*$`)
)

type spotHandler struct {
	db    *mongo.Database
	spots *Spots
}

func (h *spotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listSpotsRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getSpotRe.MatchString(r.URL.Path):
		//h.Get(w, r)
		return
	case r.Method == http.MethodPost && createSpotRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (h *spotHandler) List(w http.ResponseWriter, r *http.Request) {

	spots := []Spot{}

	spots, err := h.spots.List()
	if err != nil {
		log.Fatalf("handler: cannot fetch spots: %s", err)
	}

	jsonBytes, err := json.Marshal(spots)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

/*
func (h *spotHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getSpotRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}

	u, ok := h.store.m[matches[1]]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("spot not found"))
		return
	}
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
*/

func (h *spotHandler) Create(w http.ResponseWriter, r *http.Request) {
	var apiSpot ApiCreateSpot
	if err := json.NewDecoder(r.Body).Decode(&apiSpot); err != nil {
		internalServerError(w, r)
		return
	}

	spot, err := h.spots.CreateSpot(apiSpot.Name)
	if err != nil {
		internalServerError(w, r)
		return
	}

	jsonBytes, err := json.Marshal(spot)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
