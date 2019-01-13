package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/gowtham-munukutla/vashisht-api/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	r := mux.NewRouter()

	routes.InitRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://quiet-reef-46852.herokuapp.com"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	if os.Getenv("PORT") == "" {
		if err := http.ListenAndServe(":8000", handler); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), handler); err != nil {
			log.Fatal(err)
		}
	}
}
