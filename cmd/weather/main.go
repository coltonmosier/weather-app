package main

import (
	"fmt"
	"github.com/coltonmosier/weather-app/handlers"
)

var weatherData handlers.WeatherData

func main() {
	fmt.Println("Weather app")

	weatherData, err := handlers.QueryApi("san_antonio")
}
