package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const BaseURL = `https://api.openweathermap.org`

func main() {
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHER_API_KEY.")
		os.Exit(1)
	}
	resp, err := http.Get(fmt.Sprintf("%s/data/2.5/weather?q=London,UK&appid=%s", BaseURL, apiKey))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "unexpected response status: ", resp.StatusCode)
		os.Exit(1)
	}
	io.Copy(os.Stdout, resp.Body)
}
