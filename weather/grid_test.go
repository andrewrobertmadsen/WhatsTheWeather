package weather

import (
	"WhatsTheWeather/geocode"
	"testing"
)

func TestGridFromLatLong(t *testing.T) {
	latLong := geocode.LatLong{39.7456, -97.0892}
	got := GridFromLatLong(latLong)
	gridProperties := GridProperties{
		"TOP",
		"https://api.weather.gov/offices/TOP",
        "https://api.weather.gov/gridpoints/TOP/31,80/forecast",
        "https://api.weather.gov/gridpoints/TOP/31,80/forecast/hourly",
        "https://api.weather.gov/gridpoints/TOP/31,80",
        "https://api.weather.gov/gridpoints/TOP/31,80/stations",
	}
	want := Grid{
		"https://api.weather.gov/points/39.7456,-97.0892",
		gridProperties,
	}

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
