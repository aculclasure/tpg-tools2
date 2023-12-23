package main

import (
	"os"

	"github.com/aculclasure/weather"
)

const Usage = `Usage: weather LOCATION

Example: weather London,UK`

func main() {
	os.Exit(weather.Main())
}
