package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Conditions struct {
	Summary     string
	Temperature Temperature
}

type Temperature float64

func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp float64
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
		Summary:     resp.Weather[0].Main,
		Temperature: Temperature(resp.Main.Temp),
	}, nil
}

type Client struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.openweathermap.org",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c Client) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", c.BaseURL, location, c.APIKey)
}

func (c *Client) GetWeather(location string) (Conditions, error) {
	url := c.FormatURL(location)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return Conditions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Conditions{}, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, err
	}
	conditions, err := ParseData(data)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

func Get(location, key string) (Conditions, error) {
	c := NewClient(key)
	conditions, err := c.GetWeather(location)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

const Usage = `Usage: weather LOCATION

Example: weather London,UK`

func Main() int {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, Usage)
		return 0
	}
	apiKey := os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHER_API_KEY.")
		return 1
	}
	location := os.Args[1]
	conditions, err := Get(location, apiKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	//fmt.Println(conditions)
	fmt.Printf("%s %.1f\n", conditions.Summary, conditions.Temperature.Celsius())
	return 0
}
