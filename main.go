package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Добро пожаловать!\n(Введите help для подсказки)")

	if err := run(); err != nil {
		fmt.Printf("Ошибка: %v.", err)
		os.Exit(1)
	}
}
