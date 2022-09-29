package main

import (
	"fmt"
	"net/http"
	"strings"
)

// index handles requests to root directory of web server
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\nGo to 'localhost:8080/weather/<city name>' to see current weather in that city.\n"))
}

func main() {
	r := http.NewServeMux()
	fmt.Println("Server is listening on localhost:8080...")

	r.HandleFunc("/", index)

	// anon func writes the data for the city that is requested to it.
	r.HandleFunc("/weather/",
		func(w http.ResponseWriter, r *http.Request) {
			location := strings.SplitN(r.URL.Path, "/", 3)[2]
			data, err := QueryApi(location)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			strData := fmt.Sprintf(
				"City: %22v\nWeather: %19v\nDescription: %15v\nTemperature: %14vF\nFeels Like: %15vF\nHumidity: %17v%%\nWind Speed: %13vmph\n",
				data.City, data.Weather[0].Type, data.Weather[0].Description,
				data.Main.Temp, data.Main.Feel, data.Main.Humid, data.Wind.Speed)

			w.Write([]byte(strData))
		})

	http.ListenAndServe(":8080", r)
}
