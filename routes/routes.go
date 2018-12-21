package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vashisht-api/db"
	"github.com/vashisht-api/models"
	mgo "gopkg.in/mgo.v2"
)

var (
	DATABASE = "vashisht"
)

var dbinstance *mgo.Database

// Get list of all events
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := models.FindAllEvents()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusOK, events)
}

// Get event by id
func GetEventById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	event, err := models.FindEventById(params["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid event id")
		return
	}

	respondWithJson(w, http.StatusOK, event)
}

// Add a new event
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

	respondWithJson(w, http.StatusCreated, event)
}

// Update an existing event
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

	respondWithJson(w, http.StatusCreated, map[string]string{"result": "success"})
}

// Delete an event
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	if err := models.DeleteEvent(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, map[string]string{"result": "success"})
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	if err := models.AddUser(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, user)
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

func InitRoutes(r *mux.Router) {
	r.HandleFunc("/events", GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", GetEventById).Methods("GET")
	r.HandleFunc("/events", AddEvent).Methods("POST")
	r.HandleFunc("/events", UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", DeleteEvent).Methods("DELETE")
}
