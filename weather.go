package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CurrentWeather struct {
	Temperature float64 `json:"temperature"`
	WindSpeed   float64 `json:"windspeed"`
}

type WeatherResponse struct {
	Current CurrentWeather `json:"current_weather"`
}

func getWeather(lat, lon float64) (CurrentWeather, error) {
	weatherURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast"+
		"?latitude=%f&longitude=%f&current_weather=true", lat, lon)

	resp2, err := http.Get(weatherURL)
	if err != nil {
		fmt.Printf("Ошибка запроса данных о погоде: %v\n", err)
		return CurrentWeather{}, err
	}
	defer resp2.Body.Close()

	body, err := io.ReadAll(resp2.Body)
	if err != nil {
		fmt.Printf("Ошибка чтения данных буфера: %v\n", err)
		return CurrentWeather{}, err
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Printf("Ошибка структуры данных о погоде: %v\n", err)
		return CurrentWeather{}, err
	}

	return weather.Current, nil
}
