package weather

import "testing"

func TestGridFromLatLong(t *testing.T) {
	got := GridFromLatLong(39.7456, -97.0892)
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
