package main

import (
	"log"
	"net/http"

	"gitlab.com/gowtham-munukutla/vashisht-api/routes"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	routes.InitRoutes(r)

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
