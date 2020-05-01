package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Grid(lat, long string) {
	r, err := http.Get("https://api.weather.gov/points/" + lat + "," + long) // 39.7456,-97.0892
	if err != nil {
		panic("Error during request for grid.")
	}
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Error reading response for grid.")
	}
	/*
		type Properties struct {
			Forecast string
			ForecastHourly string
		}
		type Grid struct {
			Properties Properties
		}
		var g Grid
	*/
	var g interface{}
	err = json.Unmarshal(bodyBytes, &g)
	if err != nil {
		panic("Error parsing json for grid.")
	}
	fmt.Println(g)
}
