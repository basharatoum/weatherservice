package weatherservice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/basharatoum/weatherservice/model"
)

const (
	hot      = 80
	moderate = 60
)

func GetWeather(ctx context.Context, long float64, lat float64) (model.WeatherServiceResult, error) {
	point, err := getPointData(ctx, long, lat)
	if err != nil {
		return model.WeatherServiceResult{}, fmt.Errorf("failed to get point data: %v", err)
	}

	forecast, err := getForecastData(ctx, point)
	if err != nil {
		return model.WeatherServiceResult{}, fmt.Errorf("failed to get forecast data: %v", err)
	}

	response := model.WeatherServiceResult{
		ShortForecast: forecast.ShortForecast,
	}
	fmt.Println(forecast)

	if forecast.Temperature >= hot {
		response.Temperature = "hot"
	} else if forecast.Temperature >= moderate {
		response.Temperature = "moderate"
	} else {
		response.Temperature = "cold"
	}

	return response, nil
}

func getForecastData(ctx context.Context, point model.Point) (model.Forecast, error) {
	req, err := http.NewRequest("GET", "https://api.weather.gov/gridpoints/"+point.GridID+"/"+fmt.Sprintf("%d,%d/forecast", point.GridX, point.GridY), nil)
	if err != nil {
		return model.Forecast{}, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Accept", "application/ld+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Forecast{}, fmt.Errorf("failed to call weather API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Forecast{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var forecast model.ForecastPeriods
	if err := json.NewDecoder(resp.Body).Decode(&forecast); err != nil {
		return model.Forecast{}, fmt.Errorf("failed to decode response: %v", err)
	}

	// get most recent forecast if none labeled as "today"
	for _, period := range forecast.Periods {
		if period.Name == "Today" {
			return period, nil
		}
	}

	if len(forecast.Periods) > 0 {
		return forecast.Periods[0], nil // return the first period if no "Today" found
	}
	return model.Forecast{}, fmt.Errorf("no forecast data available")
}

func getPointData(ctx context.Context, long float64, lat float64) (model.Point, error) {

	req, err := http.NewRequest("GET", "https://api.weather.gov/points/"+fmt.Sprintf("%f,%f", lat, long), nil)
	if err != nil {
		return model.Point{}, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Accept", "application/ld+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Point{}, fmt.Errorf("failed to call weather API: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Point{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var point model.Point
	if err := json.NewDecoder(resp.Body).Decode(&point); err != nil {
		return model.Point{}, fmt.Errorf("failed to decode response: %v", err)
	}

	return point, nil
}
