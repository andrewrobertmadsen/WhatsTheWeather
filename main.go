package main

import (
	"WhatsTheWeather/geocode"
	"WhatsTheWeather/render"
	"WhatsTheWeather/weather"
	"flag"
	"fmt"
	"os"
)

func main() {
	location := flag.String("location", "Sandy, Utah", "The street address, city, state and/or zipcode of interest.")
	flag.Parse()

	latLong := geocode.LatLongFromQuery(*location)
	grid := weather.GridFromLatLong(latLong)
	forecast := weather.ForecastForGrid(grid)
	fmt.Printf("Forecast for: %q\n", *location)
	render.Render7DayForecast(forecast, os.Stdout)
}
