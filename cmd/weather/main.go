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
				`<table>
   <tr>
   <td>City</td>
   <th align="right"><strong>%v</strong></td>
   </tr>
   <tr>
   <td>Description</td>
   <td align="right"><strong>%v</strong></td>
   </tr>
   <tr>
   <td >Temperature</td>
   <td align="right"><strong>%vF</strong></td>
   </tr>
   <tr>
   <td>Feels Like</td>
   <td align="right"><strong>%vF</strong></td>
   </tr>
   <tr>
   <td>Humidity</td>
   <td align="right"><strong>%v%%</strong></td>
   </tr>
   <tr>
   <td>Wind</td>
   <td align="right"><strong>%vmph</strong></td>
   </tr>
   </table>`,
				data.City, data.Weather[0].Description,
				data.Main.Temp, data.Main.Feel, data.Main.Humid, data.Wind.Speed)

			w.Write([]byte(strData))
		})

	http.ListenAndServe(":8080", r)
}
