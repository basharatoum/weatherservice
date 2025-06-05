package model

type Point struct {
	GridID string `json:"gridId"`
	GridX  int    `json:"gridX"`
	GridY  int    `json:"gridY"`
}

type ForecastPeriods struct {
	Periods []Forecast `json:"periods"`
}

type Forecast struct {
	Name          string  `json:"name"`
	ShortForecast string  `json:"shortForecast"`
	Temperature   float64 `json:"temperature"`
}

type WeatherServiceResult struct {
	ShortForecast string `json:"shortForecast"`
	Temperature   string `json:"temperature"`
}
