package geocode

import (
	"github.com/jasonwinn/geocoder"
	"os"
)

type LatLong struct {
	Lat float64
	Long float64
}

// Provided location query string, return the geocoded latitude and longitude of that location.
func LatLongFromQuery(query string) (result LatLong) {
	key := os.Getenv("MAPQUEST_SECRET")
	geocoder.SetAPIKey(key)
	lat, long, err := geocoder.Geocode(query)
	if err != nil {
		panic(err)
	}
	return LatLong{lat, long}
}