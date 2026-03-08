package output

import (
	"fmt"

	"github.com/ongeziwe17/go-weather-cli/internal/weather"
)

func PrintWeather(result *weather.Result) {
	fmt.Println("Current Weather")
	fmt.Println("---------------")
	fmt.Printf("City: %s\n", result.City)
	fmt.Printf("Temperature: %.1f°C\n", result.Temperature)
	fmt.Printf("Condition: %s\n", result.Condition)
	fmt.Printf("Description: %s\n", result.Description)
	fmt.Printf("Humidity: %d%%\n", result.Humidity)
	fmt.Printf("Wind Speed: %.1f m/s\n", result.WindSpeed)
}
