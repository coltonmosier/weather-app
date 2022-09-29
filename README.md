# weather-app
* Website that calls a weather API from [OpenWeatherMap](https://home.openweathermap.org/) to display the current weather for a city.
* I used Golang to learn more about handling requests to and from an API.
* Also wanted to learn how to host a web server and push data to it.
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

# proof of concept
![index](https://user-images.githubusercontent.com/48069930/192921740-2ada7582-2d1c-4af5-999e-f70a5dec5551.png)
![weather](https://user-images.githubusercontent.com/48069930/192921743-6cf960c8-7581-4dcd-99c5-74e516f4421f.png)
