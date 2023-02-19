package weather_station

import "fmt"

type Observer interface {
	Update()
}

type DisplayElement interface {
	Display()
}

type CurrentConditionsDisplay struct {
	weatherData              *WeatherData
	temp, humidity, pressure float64
}

func NewCurrentConditionsDisplay(weatherData *WeatherData) *CurrentConditionsDisplay {
	observer := &CurrentConditionsDisplay{weatherData: weatherData}
	weatherData.RegisterObserver(observer)
	return observer
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Println("Current conditions:", c.temp, "F degrees and", c.humidity, "% humiditySum")
}

func (c *CurrentConditionsDisplay) Update() {
	c.temp = c.weatherData.temp
	c.humidity = c.weatherData.humidity
	c.pressure = c.weatherData.pressure
	c.Display()
}

type StatisticsDisplay struct {
	weatherData                       *WeatherData
	tempSum, humiditySum, pressureSum float64
	days                              int32
	maxTemp, minTemp                  float64
}

func NewStatisticsDisplay(weatherData *WeatherData) *StatisticsDisplay {
	observer := &StatisticsDisplay{weatherData: weatherData}
	weatherData.RegisterObserver(observer)
	observer.minTemp = 9999
	return observer
}

func (s *StatisticsDisplay) Display() {
	fmt.Println("Avg/Max/Min temperature:", s.tempSum/float64(s.days), s.maxTemp, s.minTemp)

}

func (s *StatisticsDisplay) Update() {
	s.tempSum += s.weatherData.temp
	s.humiditySum += s.weatherData.humidity
	s.pressureSum += s.weatherData.pressure
	s.days += 1
	if s.weatherData.temp > s.maxTemp {
		s.maxTemp = s.weatherData.temp
	}
	if s.weatherData.temp < s.minTemp {
		s.minTemp = s.weatherData.temp
	}
	s.Display()
}

type ForecastDisplay struct {
	weatherData              *WeatherData
	temp, humidity, pressure float64
}

func NewForecastDisplay(weatherData *WeatherData) *ForecastDisplay {
	observer := &ForecastDisplay{weatherData: weatherData}
	weatherData.RegisterObserver(observer)
	return observer
}

func (f *ForecastDisplay) Display() {
	fmt.Println("Forecast:", f.temp, "F degrees and", f.humidity, "% humiditySum and", f.pressure, "kPa pressureSum")

}

func (f *ForecastDisplay) Update() {
	f.temp = f.weatherData.temp
	f.humidity = f.weatherData.humidity
	f.pressure = f.weatherData.pressure
	f.Display()
}

type ThirdPartyDisplay struct {
	weatherData              *WeatherData
	temp, humidity, pressure float64
}

func NewThirdPartyDisplay(weatherData *WeatherData) *ThirdPartyDisplay {
	display := ThirdPartyDisplay{weatherData: weatherData}
	weatherData.RegisterObserver(&display)
	return &display
}

func (t *ThirdPartyDisplay) Update() {
	t.temp = t.weatherData.temp
	t.humidity = t.weatherData.humidity
	t.pressure = t.weatherData.pressure
	t.Display()
}

func (t *ThirdPartyDisplay) Display() {
	fmt.Println("ThirdParty:", t.temp, "F degrees and", t.humidity, "% humiditySum and", t.pressure, "kPa pressureSum")

}

type HeatIndexDisplay struct {
	WeatherData *WeatherData
	temp        float64
	humidity    float64
	pressure    float64
}

func (h *HeatIndexDisplay) Display() {
	fmt.Println("Heat index is", h.heatIndex())
}

func (h *HeatIndexDisplay) Update() {
	h.temp = h.WeatherData.temp
	h.humidity = h.WeatherData.humidity
	h.pressure = h.WeatherData.pressure

	h.Display()
}

func (h *HeatIndexDisplay) heatIndex() float64 {
	first := 16.923 + 1.85212e-1*h.temp + 5.37941*h.humidity
	quad := -1.00254e-1*h.temp*h.humidity + 9.41695e-3 + h.temp*h.temp + 7.28898e-3*h.humidity*h.humidity
	cubic := 3.45372e-4*h.temp*h.temp*h.humidity - 8.14971e-4*h.temp*h.humidity*h.humidity - 3.8646e-5*h.temp*h.temp*h.temp + 2.91583e-5*h.humidity*h.humidity*h.humidity
	quar := 1.02102e-5 * h.temp * h.temp * h.humidity * h.humidity
	return first + quad + cubic + quar
}

func NewHeatIndexDisplay(weatherData *WeatherData) *HeatIndexDisplay {
	observer := &HeatIndexDisplay{WeatherData: weatherData}
	weatherData.RegisterObserver(observer)
	return observer
}
