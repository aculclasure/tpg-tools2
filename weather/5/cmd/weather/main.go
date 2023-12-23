package main

import (
	"fmt"
	"os"

	"github.com/aculclasure/weather"
)

const Usage = `Usage: weather LOCATION

Example: weather London,UK`

func main() {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHER_API_KEY.")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, Usage)
		os.Exit(0)
	}
	location := os.Args[1]
	conditions, err := weather.Get(location, apiKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(conditions)
}
