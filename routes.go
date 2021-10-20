package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		index,
	},
	Route{
		"getAllCountry",
		"GET",
		"/api/country",
		getAllCountry,
	},
	Route{
		"GetAllCity",
		"GET",
		"/api/city/{country}",
		getAllCities,
	},
	Route{
		"GetWeather",
		"GET",
		"/api/weather/{city}",
		getWeather,
	},
}
