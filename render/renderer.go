package render

import (
	"WhatsTheWeather/weather"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io"
)

func Render7DayForecast(forecast weather.Forecast, writer io.Writer) {
	forecastNames := []string{}
	forecastShorts := []string{}
	forecastTemps := []string{}
	for _, period := range forecast.Properties.Periods {
		forecastNames = append(forecastNames, period.Name)
		forecastShorts = append(forecastShorts, period.ShortForecast)
		forecastTemps = append(forecastTemps, fmt.Sprintf("%d°%v", period.Temperature, period.TemperatureUnit))
	}
	table := tablewriter.NewWriter(writer)
	table.SetHeader(forecastNames)
	table.Append(forecastShorts)
	table.Append(forecastTemps)
	table.Render()
}
