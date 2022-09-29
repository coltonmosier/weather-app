# weather-app
* Website that calls a weather API from [OpenWeatherMap](https://home.openweathermap.org/) to display the current weather for a city.
* I used Golang to learn more about making requests from an API and how to store and push the data to a web server.
# Usage
1. Get the application:  
`git clone https://github.com/coltonmosier/weather-app`

2. Create a free account on [OpenWeatherMap](https://home.openweathermap.org/), which will email you your free ApiKey
3. Add your ApiKey to the .apiConfig file within weather-app/
4. Run application  
`go run weather-app/cmd/weather/*.go`
5. Visit 'localhost:8080' and follow instructions.