package main

import (
	"fmt"
	"io"
	"net/http"
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
	url := weather.FormatURL(weather.BaseURL, location, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "unexpected response status: ", resp.StatusCode)
		os.Exit(1)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	cond, err := weather.ParseData(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(cond)
}
