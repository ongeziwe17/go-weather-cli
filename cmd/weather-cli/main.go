package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./cmd/weather-cli \"Cape Town\"")
		os.Exit(1)
	}
	city := strings.Join(os.Args[1:], " ")
	fmt.Printf("Fetching weather for: %s\n", city)
}
