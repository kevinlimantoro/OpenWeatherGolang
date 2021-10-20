package config

import (
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Config Config `json:"config"`
}

type Config struct {
	Dev  EnvConfig `json:"dev"`
	Prod EnvConfig `json:"prod"`
}

type EnvConfig struct {
	OpenWeather OpenWeatherConfig `json:"OpenWeather"`
}

type OpenWeatherConfig struct {
	APIKey      string `json:"APIKey"`
	APIEndpoint string `json:"APIEndpoint"`
}

func GetAPIEndpoint(env, city string) string {
	configuration := Configuration{}
	gonfig.GetConf("./config.json", &configuration)
	if env == "prod" {
		res := strings.ReplaceAll(configuration.Config.Prod.OpenWeather.APIEndpoint,
			"{API_KEY}",
			configuration.Config.Prod.OpenWeather.APIKey)
		return strings.ReplaceAll(res, "{CITY}", city)
	} else {
		res := strings.ReplaceAll(configuration.Config.Dev.OpenWeather.APIEndpoint,
			"{API_KEY}",
			configuration.Config.Dev.OpenWeather.APIKey)
		return strings.ReplaceAll(res, "{CITY}", city)
	}
}
