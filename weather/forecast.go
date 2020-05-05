package weather

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Forecast struct {
	Properties ForecastProperties
}

type ForecastProperties struct {
	Updated string
	Units string
	GeneratedAt string
	UpdateTime string
	ValidTimes string
	Periods []Period
}

type Period struct {
	Number int
	Name string
	StartTime string
	EndTime string
	IsDaytime bool
	Temperature int
	TemperatureUnit string
	TempuratureTrend string
	WindSpeed string
	WindDirection string
	Icon string
	ShortForcast string
	DetailedForcast string
}

func ForecastForGrid(grid Grid) Forecast {
	r, err := http.Get(grid.Properties.Forecast)
	if err != nil {
		panic("Error during request for forecast.")
	}
	defer r.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Error reading response for forecast.")
	}
	var forecast Forecast
	err = json.Unmarshal(bodyBytes, &grid)
	if err != nil {
		panic("Error parsing json for forecast.")
	}
	return forecast
}
