package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/gowtham-munukutla/vashisht-api/routes"

	"github.com/gorilla/mux"

	"gitlab.com/gowtham-munukutla/vashisht-api/mailer"
	"gitlab.com/gowtham-munukutla/vashisht-api/models"
)

func main() {

	user := &models.User{
		Name:         "Gowtham Munukutla",
		EmailAddress: "gowtham.m81197@gmail.com",
	}

	err := mailer.SendRegistrationEmail(user)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Mail sent")
	}

	r := mux.NewRouter()

	routes.InitRoutes(r)

	if os.Getenv("PORT") == "" {
		if err := http.ListenAndServe(":8000", r); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
			log.Fatal(err)
		}
	}
}
