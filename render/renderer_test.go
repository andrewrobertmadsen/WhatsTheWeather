package render

import (
	"WhatsTheWeather/weather"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func loadForecast() *weather.Forecast {
	cwd, _ := os.Getwd()
	fileBytes, err := ioutil.ReadFile(path.Join(cwd, "resources", "forecast.json"))
	var f weather.Forecast
	err = json.Unmarshal(fileBytes, &f)
	if err != nil {
		panic(err)
	}
	return &f
}

func TestRender7DayForecast(t *testing.T) {
	f := loadForecast()
	buf := bytes.Buffer{}
	Render7DayForecast(*f, &buf)
	got := buf.String()
	want := "Hi"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
