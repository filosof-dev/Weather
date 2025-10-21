package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Country   string  `json:"country"`
	} `json:"results"`
}

var ErrReqGeoData = errors.New("этап запроса геоданных")
var ErrStructGeoData = errors.New("этап заполнения структуры геоданных")
var ErrNotFoundCity = errors.New("отсутствуют данные города в структуре")

func getCoordinates(city string) (lat, lon float64, name string, err error) {
	geoURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", city)

	resp, err := http.Get(geoURL)
	if err != nil {
		fmt.Println("Ошибка запроса геоданных.")                      // Текст для пользователя.
		return 0, 0, "", fmt.Errorf("функция run: %w", ErrReqGeoData) // Передача ошибки выше для возможной обработки, например лог.
	}
	defer resp.Body.Close()

	bodyGeo, _ := io.ReadAll(resp.Body)

	var geo GeoResponse
	if err := json.Unmarshal(bodyGeo, &geo); err != nil {
		fmt.Println("Ошибка структуры геоданных.")
		return 0, 0, "", fmt.Errorf("функция run: %w", ErrStructGeoData)
	}

	if len(geo.Results) == 0 {
		fmt.Println("Город не найден, попробуйте ввести другой город.")
		return 0, 0, "", fmt.Errorf("функция run: %w", ErrNotFoundCity)
	}

	return geo.Results[0].Latitude, geo.Results[0].Longitude, geo.Results[0].Name, nil
}
