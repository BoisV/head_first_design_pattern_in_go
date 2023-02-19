package weather_station

type Subject interface {
	RegisterObserver(observer *Observer)
	RemoveObserver(observer *Observer)
	NotifyObservers()
}

type WeatherData struct {
	observers                []Observer
	temp, humidity, pressure float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{observers: []Observer{}}
}

func (w *WeatherData) SetMeasurements(temp, humidity, pressure float64) {
	w.temp = temp
	w.humidity = humidity
	w.pressure = pressure
	w.measurementsChanged()
}

func (w *WeatherData) RegisterObserver(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherData) RemoveObserver(observer interface{}) {
	for i := 0; i < len(w.observers); i++ {
		if w.observers[i] == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
		}
	}
}

func (w *WeatherData) NotifyObservers() {
	for i := 0; i < len(w.observers); i++ {
		w.observers[i].Update()
	}
}

func (w *WeatherData) measurementsChanged() {
	w.NotifyObservers()
}
