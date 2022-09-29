package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

// WeatherData stores the json that is received from the API request
type WeatherData struct {
	City    string `json:"name"`
	Weather []struct {
		Type        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp  float64 `json:"temp"`
		Feel  float64 `json:"feels_like"`
		Humid int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	}
}

// ApiConfig is a struct to store the key from the .apiConfig file
type apiConfig struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

// ApiSetUp reads from .apiConfig file to grab and store the Api Key into
// ApiConfig struct.
func ApiSetUp(fileName string) (apiConfig, error) {
	var setUpApi apiConfig
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return apiConfig{}, errors.New("could not read .apiConfig file")
	}
	if len(bytes) == 0 {
		return apiConfig{}, errors.New("no key loaded")
	}

	err = json.Unmarshal(bytes, &setUpApi)
	if err != nil {
		return apiConfig{}, errors.New("could not get json from .apiConfig file")
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
