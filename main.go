package main

import (
	"encoding/json"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"

	"github.com/vashisht-api/db"
	"github.com/vashisht-api/models"
)

var (
	DATABASE   = "vashisht"
	COLLECTION = "events"
)

var dbinstance *mgo.Database

// Get list of all events
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := models.FindAllEvents(dbinstance)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, events)
}

// Get event by id
func GetEventById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	event, err := models.FindEventById(params["id"], dbinstance)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid event id")
		return
	}

	respondWithJson(w, http.StatusOK, event)
}

func respondWithError(w http.ResponseWriter, httpCode int, message string) {
	respondWithJson(w, httpCode, map[string]string{"error": message})
}

func respondWithJson(w http.ResponseWriter, httpCode int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}

func init() {
	dbinstance = db.GetDbInstance()
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/events", GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", GetEventById).Methods("GET")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
