package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type WeatherData struct {
	City    string `json:"city"`
	Weather string `json:"weather"`
}

// ApiConfig is a struct to store the key from the .apiConfig file
type ApiConfig struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

// ApiSetUp reads from .apiConfig file to grab and store the Api Key into
// ApiConfig struct.
func ApiSetUp(fileName string) (ApiConfig, error) {
	var setUpApi ApiConfig
	fmt.Println("api config")
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return ApiConfig{}, errors.New("Could not read .apiConfig file")
	}
	if len(bytes) == 0 {
		return ApiConfig{}, errors.New("No key loaded")
	}

	err = json.Unmarshal(bytes, &setUpApi)
	if err != nil {
		return ApiConfig{}, errors.New("Could not get json from .apiConfig file")
	}

	return setUpApi, nil
}

// QueryApi sends a query about the city passed in as a string variable
// to the API and captures the response.
func QueryApi(city string) (WeatherData, error) {
	apiKey, err := ApiSetUp(".apiConfig")
	var data WeatherData
	if len(apiKey.OpenWeatherMapApiKey) == 0 {
		return WeatherData{}, errors.New("no api key loaded")
	}
	if err != nil {
		return WeatherData{}, err
	}

	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + apiKey.OpenWeatherMapApiKey + "&units=imperial"

	res, err := http.Get(url)
	if err != nil {
		return WeatherData{}, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return WeatherData{}, err
	}

	return data, nil
}
