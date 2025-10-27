# Weather — консольный погодный сервис на Go.  
### Небольшое консольное приложение, которое получает текущую погоду по названию города, используя публичное API Open-Meteo.  

## Возможности:  
- Поиск города по названию через API геокодинга.  
- Получение текущей температуры и скорости ветра.  
- Проверка корректности ввода (латиница, пробелы, дефисы).  
- Простые команды управления: help, exit.  
- Обработка типовых ошибок: отсутствие соединения, неверный город, сбой при разборе JSON.

## Стек технологий:  
- Go (1.24.4)  
- net/http, net/url, encoding/json, bufio, unicode, errors, io, os, strings, fmt  
- Публичное API:  
  - Open-Meteo Geocoding API  
  - Open-Meteo Weather API

## Структура проекта:  
Weather/  
- main.go         - *Точка входа*  
- run.go          - *Основной цикл программы, парсинг команд*   
- geo.go          - *Запрос геокоординат по названию города*  
- weather.go      - *Запрос текущей погоды по координатам*  
- utils.go        - *Вспомогательные проверки и функции*  
- go.mod          - *Зависимости*

## Установка и запуск:
### Клонировать репозиторий:  
`git clone https://github.com/filosof-dev/Weather.git`  
`cd Weather`  

### Установить зависимости:  
`go mod tidy`  

### Запуск приложения:  
`go run .`  

### Сборка исполняемого файла для Windows и его запуск:  
`go build -o weather.exe`  
`.\weather`  

## Пример работы:  
Запуск приложения и ввод города:  
<img width="363" height="86" alt="image" src="https://github.com/user-attachments/assets/d261ebed-1b53-4d2b-9bd9-1dc07573b4db" />  

Результат:  
<img width="403" height="65" alt="image" src="https://github.com/user-attachments/assets/a6b7119a-ad6a-42c4-b130-e835b99b6d58" />   

Работа команда help:  
<img width="391" height="109" alt="image" src="https://github.com/user-attachments/assets/ef68569f-5a50-484d-b4ad-23b570e50917" />   

Выход из приложения:  
<img width="386" height="70" alt="image" src="https://github.com/user-attachments/assets/d22c4a8c-36ed-4755-8fe6-f25e8652eaa1" />  

## Лицензия:  
Проект распространяется под лицензией [MIT](./LICENSE)  
Это свободная лицензия, позволяющая использовать, изменять и распространять код без ограничений, но с упоминанием автора.  

