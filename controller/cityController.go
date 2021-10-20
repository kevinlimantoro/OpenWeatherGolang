package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kevinlimantoro/WeatherAPI/cache"
	"github.com/kevinlimantoro/WeatherAPI/config"
)

const cacheKey = "CountryCity"

func InitCache() {
	cache.InitCache()
}

func getCountryCityFromJson() map[string][]string {
	countryCity := make(map[string][]string)
	c, err := cache.MyCache.Get(cacheKey)
	if err != nil {
		log.Fatal(err)
	} else if c != nil {
		fmt.Println("Hit cache")
		json.Unmarshal(c, &countryCity)
		return countryCity
	}
	jsonFile, err := os.Open("./countries.min.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	fmt.Println("Successfully Opened countries.min.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &countryCity)
	cache.MyCache.Set(cacheKey, countryCity, 24*time.Hour)
	return countryCity
}

func GetCountryCities(param string) []string {
	countryCity := getCountryCityFromJson()
	if param == "" {
		countries := make([]string, 0)
		for k, _ := range countryCity {
			countries = append(countries, k)
		}
		return countries
	} else {
		return countryCity[param]
	}
}

func GetWeather(city string) interface{} {
	var weather Weather
	w, err := cache.MyCache.Get(cacheKey + city)
	if err != nil {
		log.Fatal(err)
	} else if w != nil {
		fmt.Println("Hit cache")
		json.Unmarshal(w, &weather)
		return weather
	}
	resp, err := http.Get(config.GetAPIEndpoint("prod", city))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &weather)
	cache.MyCache.Set(cacheKey+city, weather, 3*time.Minute)
	return weather
}
