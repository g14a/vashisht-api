package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/gowtham-munukutla/vashisht-api/routes"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	routes.InitRoutes(r)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), r); err != nil {
		log.Fatal(err)
	}
}
