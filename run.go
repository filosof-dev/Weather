package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func run() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите город или команду, например Moscow: ")

		input, err := reader.ReadString('\n') // Считывание данных.
		if err != nil {
			fmt.Printf("Ошибка чтения ввода, попробуйте ещё раз. %v", err)
			continue
		}

		// --> Работа с введенным текстом, очистка от артефактов.
		input = strings.TrimSpace(input)                 // Удаляет /n/r с краев.
		input = strings.Join(strings.Fields(input), " ") // Схлопывает лишние пробелы в один.

		if !isLatinOnly(input) { // Проверка на латиницу.
			fmt.Println("Внимание! Название города должно быть на английском языке латинскими буквами.")
			continue
		}

		input = url.QueryEscape(input) // Формализует url. Вместо пробела ставит +.
		// <--

		switch input {
		case "exit":
			fmt.Println("До свидания!")
			return nil
		case "help":
			fmt.Println("Вводите город на английском языке, например Moscow.\n" +
				"Доступные команды:\n" +
				"help - помощь по работе с программой.\n" +
				"exit - выход из программы.")
			continue
		}

		lat, lon, name, err := getCoordinates(input)
		if err != nil { // Проверки на конкретные ошибки, можно расширить функционал, например запись логов.
			if errors.Is(err, ErrReqGeoData) {
				continue
			} else if errors.Is(err, ErrStructGeoData) {
				continue
			} else if errors.Is(err, ErrNotFoundCity) {
				continue
			}

			return err
		}

		weather, err := getWeather(lat, lon)
		if err != nil {
			return err
		}

		fmt.Printf("Температура в %s: %.1f°C, ветер: %.1f м/с.\n", name, weather.Temperature, weather.WindSpeed)
	}
}
