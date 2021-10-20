package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinlimantoro/WeatherAPI/controller"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is Cool!")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Println("Endpoint Hit: index")
}

func getAllCountry(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllCountry")
	cities := controller.GetCountryCities("")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cities); err != nil {
		panic(err)
	}
}

func getAllCities(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getAllCities")
	vars := mux.Vars(r)
	country := vars["country"]
	cities := controller.GetCountryCities(country)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cities); err != nil {
		panic(err)
	}
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getWeather")
	vars := mux.Vars(r)
	city := vars["city"]
	res := controller.GetWeather(city)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}
