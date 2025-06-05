package model

type Point struct {
	ForecastOffice string `json:"forecastOffice"`
	GridID         string `json:"gridId"`
	GridX          int    `json:"gridX"`
	GridY          int    `json:"gridY"`
}

type ForecastPeriods struct {
	Periods []Forecast `json:"periods"`
}

type Forecast struct {
	Name          string `json:"name"`
	ShortForecast string `json:"shortForecast"`
	Temperature   struct {
		Value float64 `json:"value"`
	} `json:"temperature"`
}

type WeatherServiceResult struct {
	ShortForecast string `json:"shortForecast"`
	Temperature   string `json:"temperature"`
}
