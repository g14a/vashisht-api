package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.com/gowtham-munukutla/vashisht-api/db"
	"gitlab.com/gowtham-munukutla/vashisht-api/models"
	mgo "gopkg.in/mgo.v2"
)

// This is the db of the fest
var (
	DATABASE = "vashisht"
)

var dbinstance *mgo.Database

// GetAllEvents returns all the events in the db. Refer to models package for more info
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := models.FindAllEvents()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, events)
}

// GetEventByID returns an event by id
func GetEventByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	event, err := models.FindEventByID(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid event id")
		return
	}

	respondWithJSON(w, http.StatusOK, event)
}

// AddEvent adds a new event to the db
func AddEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	if err := models.AddEvent(&event); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, event)
}

// UpdateEvent updates an event through PUT
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	if err := models.UpdateEvent(&event); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

// DeleteEvent deletes an event from the db
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	if err := models.DeleteEvent(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

// AddUser adds a new user
func AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	if err := models.AddUser(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

// GetAllUsers returns all the users registered for the fest
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	users, err := models.GetAllUsers()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

func respondWithError(w http.ResponseWriter, httpCode int, message string) {
	respondWithJSON(w, httpCode, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, httpCode int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(response)
}

func init() {
	dbinstance = db.GetDbInstance()
}

// InitRoutes initializes all the http routes ...
func InitRoutes(r *mux.Router) {
	r.HandleFunc("/events", GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", GetEventByID).Methods("GET")
	r.HandleFunc("/events", AddEvent).Methods("POST")
	r.HandleFunc("/events", UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", DeleteEvent).Methods("DELETE")

	// User routes
	r.HandleFunc("/users", AddUser).Methods("POST")
	r.HandleFunc("/users", GetAllUsers).Methods("GET")
}
