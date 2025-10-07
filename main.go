package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GeoResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Country   string  `json:"country"`
	} `json:"results"`
}

type CurrentWeather struct {
	Temperature float64 `json:"temperature"`
	WindSpeed   float64 `json:"windspeed"`
}

type WeatherResponse struct {
	Current CurrentWeather `json:"current_weather"`
}

func main() {
	fmt.Print("Введите город: ")
	var city string
	fmt.Fscan(os.Stdin, &city)

	geoURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", city)

	resp, err := http.Get(geoURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyGeo, _ := io.ReadAll(resp.Body)

	var geo GeoResponse
	if err := json.Unmarshal(bodyGeo, &geo); err != nil {
		panic(err)
	}

	if len(geo.Results) == 0 {
		fmt.Println("Город не найден")
		return
	}

	lat := geo.Results[0].Latitude
	lon := geo.Results[0].Longitude

	weatherURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast"+
		"?latitude=%f&longitude=%f&current_weather=true", lat, lon)

	resp2, err := http.Get(weatherURL)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	body, err := io.ReadAll(resp2.Body)
	if err != nil {
		panic(err)
	}

	var weather WeatherResponse
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	/*
		fmt.Printf("Температура в %s: %.1f°C",
			geo.Results[0].Name,
			weather.Current.Temperature)
	*/

}

//https://api.open-meteo.com/v1/forecast" +
//		"?latitude=55.75&longitude=37.61&current_weather=true
