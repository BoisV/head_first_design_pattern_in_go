package weather_station

import (
	"testing"
)

func TestWeatherStation(t *testing.T) {
	weatherData := NewWeatherData()
	_ = NewCurrentConditionsDisplay(weatherData)
	_ = NewForecastDisplay(weatherData)
	_ = NewStatisticsDisplay(weatherData)

	weatherData.SetMeasurements(80, 65, 30.4)
	println()
	weatherData.SetMeasurements(82, 72, 29.2)
	println()
	weatherData.SetMeasurements(78, 90, 29.2)

}
