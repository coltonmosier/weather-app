package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

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
func QueryApi(city string) ([]byte, error) {
	apiKey, err := ApiSetUp(".apiConfig")
	var data []byte
	if len(apiKey.OpenWeatherMapApiKey) == 0 {
		return data, errors.New("no api key loaded")
	}
	if err != nil {
		return data, err
	}
	city = "blah"
	// https://api.openweathermap.org/data/2.5/weather?q={ city }&appid={API key}&units=imperial

	return data, nil
}
