package geocode

import "github.com/jasonwinn/geocoder"

type LatLong struct {
	Lat float64
	Long float64
}

// Provided location query string, return the geocoded latitude and longitude of that location.
func LatLongFromQuery(query string) (result LatLong) {
	lat, long, err := geocoder.Geocode(query)
	if err != nil {
		panic("Error during geocoding of address query.")
	}
	return LatLong{lat, long}
}