package models

type period struct {
	number int
	name string
	startTime string
	endTime string
	isDaytime bool
	temperature int
	temperatureUnit string
	tempuratureTrend string
	windSpeed string
	windDirection string
	icon string
	shortForcast string
	detailedForcast string
}