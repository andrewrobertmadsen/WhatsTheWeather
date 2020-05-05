package weather

import (
	"reflect"
	"testing"
)

func TestForecastForGrid(t *testing.T) {
	gridProperties := GridProperties{
		"TOP",
		"https://api.weather.gov/offices/TOP",
		"https://api.weather.gov/gridpoints/TOP/31,80/forecast",
		"https://api.weather.gov/gridpoints/TOP/31,80/forecast/hourly",
		"https://api.weather.gov/gridpoints/TOP/31,80",
		"https://api.weather.gov/gridpoints/TOP/31,80/stations",
	}
	grid := Grid{
		"https://api.weather.gov/points/39.7456,-97.0892",
		gridProperties,
	}
	got := ForecastForGrid(grid)
	want := Forecast{ForecastProperties{
		"",
		"",
		"",
		"",
		"",
		[]Period{
			{
				0,
				"",
				"",
				"",
				true,
				0,
				"",
				"",
				"",
				"",
				"",
				"",
				"",
			},
		},
	}}

	if reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}