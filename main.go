package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/gowtham-munukutla/vashisht-api/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	routes.InitRoutes(r)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	if os.Getenv("PORT") == "" {
		if err := http.ListenAndServe(":5000", handlers.CORS(headers, methods, origins)(r)); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := http.ListenAndServe(":"+os.Getenv("PORT"), handlers.CORS(headers, methods, origins)(r)); err != nil {
			log.Fatal(err)
		}
	}
}
