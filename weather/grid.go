package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Grid struct {
	Id string
	Properties GridProperties
}

type GridProperties struct {
	CWA string
	ForecastOffice string
	Forecast string
	ForecastHourly string
	ForecastGridData string
	ObservationStations string
}

func GridFromLatLong(lat, long float64) Grid {
	endpoint := fmt.Sprintf("https://api.weather.gov/points/%+v,%+v", lat, long)
	r, err := http.Get(endpoint)
	if err != nil {
		panic("Error during request for grid.")
	}
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Error reading response for grid.")
	}
	var grid Grid
	err = json.Unmarshal(bodyBytes, &grid)
	if err != nil {
		panic("Error parsing json for grid.")
	}
	return grid
}
