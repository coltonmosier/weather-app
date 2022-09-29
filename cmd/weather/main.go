package main

import (
	"fmt"
	"github.com/coltonmosier/weather-app/handlers"
	"net/http"
)

var weatherData handlers.WeatherData

func main() {
	fmt.Println("Weather app")

	weather := handlers.Weather()
	home := handlers.Home()

	r := http.NewServeMux()
	r.Handle("/", home)
	r.HandleFunc("/weather/", weather)

	http.ListenAndServe(":8080", r)
}
