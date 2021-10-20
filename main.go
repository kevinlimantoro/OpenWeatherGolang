package main

import (
	"log"
	"net/http"

	"github.com/kevinlimantoro/WeatherAPI/controller"
)

func main() {
	controller.InitCache()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
