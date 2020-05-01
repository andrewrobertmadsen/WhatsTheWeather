package models

import "testing"

func TestLatLongFromQuery(t *testing.T) {
	query := "Sandy, Utah"
	got := LatLongFromQuery(query)
	want := LatLong{40.572851, -111.83345}

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}