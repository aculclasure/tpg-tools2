package weather

import (
	"encoding/json"
	"fmt"
)

const BaseURL = `https://api.openweathermap.org`

type Conditions struct {
	Summary string
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
}

func ParseData(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf(
			"got error '%s' trying to parse openweather response: %s", err, string(data))
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf(
			"got error '%s' trying to parse openweather response: %s, want at least one weather element", err, string(data))
	}
	return Conditions{
		Summary: resp.Weather[0].Main,
	}, nil
}

func FormatURL(baseURL, location, key string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", baseURL, location, key)
}
