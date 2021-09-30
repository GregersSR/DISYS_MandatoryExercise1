package main

import (
	"log"
	"net/http"

	"github.com/gregerssr/DISYS_MandatoryExercise1/rest/api"
	"github.com/gregerssr/DISYS_MandatoryExercise1/rest/storage"
)

func main() {
	log.Printf("Server started")
	storage.Init()

	router := api.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
