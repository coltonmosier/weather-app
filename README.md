# weather-app
* Website that calls a weather API from [OpenWeatherMap](https://home.openweathermap.org/) to display the current weather for a city.
* I used Golang to learn more about handling requests to and from an API.
* Also wanted to begin learning how to host a web server and push data to it.
# usage
1. Get the application:  
`git clone https://github.com/coltonmosier/weather-app`

2. Create a free account on [OpenWeatherMap](https://home.openweathermap.org/), which will email you your free ApiKey
3. Add your ApiKey to the .apiConfig file within weather-app/
4. Build and run application  
`cd weather-app/`  
`go build -o weather ./cmd/weather/`  
`./weather`
5. Visit 'localhost:8080' and follow instructions.

# POC
![index](https://user-images.githubusercontent.com/48069930/192920817-acd14b45-ce79-4fe4-b5a2-50439e23c912.png)
![weather](https://user-images.githubusercontent.com/48069930/192920853-ef5ba9ff-3e48-48e9-913e-fbc7ca59477a.png)
