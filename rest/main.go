package main

import (
	"log"
	"net/http"

	"github.com/gregerssr/DISYS_MandatoryExercise1/api"
)

func main() {
	log.Printf("Server started")

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
