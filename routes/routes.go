package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"gitlab.com/gowtham-munukutla/vashisht-api/mailer"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gitlab.com/gowtham-munukutla/vashisht-api/models"
)

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

	id, _ := strconv.Atoi(params["id"])

	event, err := models.FindEventByID(id)

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
		log.Println("Error:", err)
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	if err := models.AddUser(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, user)

	// Send the mail after the users is added into the db
	mailErr := mailer.SendRegistrationEmail(&user)

	log.Println(mailErr)
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

// AddRegistration adds a registration
func AddRegistration(w http.ResponseWriter, r *http.Request) {
	var registration models.Registration

	params := mux.Vars(r)
	userID := params["userid"]
	eventID := params["eventid"]

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Println(err.Error())
	}

	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		log.Println(err.Error())
	}

	registration.EventID = eventIDInt
	registration.UserID = userIDInt
	registration.RegID = uuid.NewV4().String()

	if err := models.AddRegistration(&registration); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, registration)
}

// CancelRegistration cancels a registration whenever a user wants to cancel it
func CancelRegistration(w http.ResponseWriter, r *http.Request) {
	var registration models.Registration

	params := mux.Vars(r)

	userID := params["userid"]
	eventID := params["eventid"]

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		log.Println(err.Error())
	}

	eventIDInt, err := strconv.Atoi(eventID)
	if err != nil {
		log.Println(err.Error())
	}

	registration.EventID = eventIDInt
	registration.UserID = userIDInt

	ok, err := models.CancelRegistration(registration)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": strconv.FormatBool(ok)})

}

// GetEventsOfUsers gets all the events of the users
func GetEventsOfUsers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	params := mux.Vars(r)
	userID, _ := strconv.Atoi(params["userid"])

	events, err := models.GetEventsOfUser(userID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, &events)
}

// GetUsersForEvent returns all users for an event
func GetUsersForEvent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	eventID, _ := strconv.Atoi(params["eventid"])

	users, err := models.GetUsersForEvent(eventID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, &users)
}

// CheckIfUserRegisteredForEvent checks if a user registered for an event
func CheckIfUserRegisteredForEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, _ := strconv.Atoi(params["userid"])
	eventID, _ := strconv.Atoi(params["eventid"])

	ok, err := models.CheckIfUserRegisteredForEvent(userID, eventID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ok)
}

// CheckIfUserRegisteredForEventByMongoID checks if a user registered for an event using MongoID
func CheckIfUserRegisteredForEventByMongoID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	mongoIDStr := params["mongoid"]
	eventID, _ := strconv.Atoi(params["eventid"])

	ok, err := models.CheckIfUserRegisteredForEventByMongoID(mongoIDStr, eventID)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ok)
}

// Login logs the user into the application
func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseUser, err := models.Login(user.EmailAddress, user.Password)

	if err != nil {
		respondWithJSON(w, http.StatusNotFound , map[string]string{"error": "user not found"})
		return
	}

	if responseUser == nil {
		respondWithJSON(w, http.StatusNotFound, map[string]string{"error": "user not found"})
		return
	}

	respondWithJSON(w, http.StatusOK, &responseUser)
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

// InitRoutes initializes all the http routes ...
func InitRoutes(r *mux.Router) {

	r.HandleFunc("/events", GetAllEvents).Methods("GET")
	r.HandleFunc("/events/{id}", GetEventByID).Methods("GET")
	r.HandleFunc("/events", AddEvent).Methods("POST")
	r.HandleFunc("/events", UpdateEvent).Methods("PUT")
	r.HandleFunc("/events/{id}", DeleteEvent).Methods("DELETE")
	r.HandleFunc("/events/{eventid}/users", GetUsersForEvent).Methods("GET")

	// User routes
	r.HandleFunc("/users", AddUser).Methods("POST")
	r.HandleFunc("/users", GetAllUsers).Methods("GET")
	r.HandleFunc("/users/login", Login).Methods("POST")
	r.HandleFunc("/users/{userid}/events", GetEventsOfUsers).Methods("GET")
	r.HandleFunc("/users/{userid}/events/{eventid}/register", AddRegistration).Methods("POST")
	r.HandleFunc("/users/{userid}/events/{eventid}/cancel", CancelRegistration).Methods("DELETE")
	r.HandleFunc("/users/{userid}/events/{eventid}/check", CheckIfUserRegisteredForEvent).Methods("GET")
	r.HandleFunc("/users/{mongoid}/events/{eventid}/checkMongoID", CheckIfUserRegisteredForEventByMongoID).Methods("GET")
}
