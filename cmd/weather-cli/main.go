package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ongeziwe17/go-weather-cli/internal/output"
	"github.com/ongeziwe17/go-weather-cli/internal/weather"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found")
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./cmd/weather-cli \"Cape Town\"")
		os.Exit(1)
	}

	city := strings.Join(os.Args[1:], " ")

	client := weather.NewClient()
	service := weather.NewService(client)

	result, err := service.GetCurrentWeather(city)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output.PrintWeather(result)
}
