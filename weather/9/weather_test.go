package weather_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/aculclasure/weather"
	"github.com/google/go-cmp/cmp"
)

func TestParseData_CorrectlyParsesJSONData(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/valid_weather_response.json")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		City:        "London",
		CountryCode: "GB",
		Summary:     "Clouds",
		Temperature: 283.38,
	}
	got, err := weather.ParseData(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseData_ReturnsErrorOnEmptyData(t *testing.T) {
	t.Parallel()
	_, err := weather.ParseData([]byte{})
	if err == nil {
		t.Fatal("want error parsing empty response, got nil")
	}
}

func TestParseData_ReturnsErrorOnIncompleteJSONData(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/incomplete_weather_response.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = weather.ParseData(data)
	if err == nil {
		t.Fatal("want error parsing incomplete response, got nil")
	}
}

func TestFormatURL_ReturnsCorrectURLForGivenInputs(t *testing.T) {
	t.Parallel()
	c := weather.NewClient("dummyAPIKey")
	want := "https://api.openweathermap.org/data/2.5/weather?q=London,UK&appid=dummyAPIKey"
	got := c.FormatURL("London,UK")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetWeather_ReturnsExpectedConditions(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "testdata/valid_weather_response.json")
	}))
	defer ts.Close()
	c := weather.NewClient("dummyAPIKey")
	c.BaseURL = ts.URL
	c.HTTPClient = ts.Client()
	want := weather.Conditions{
		City:        "London",
		CountryCode: "GB",
		Summary:     "Clouds",
		Temperature: 283.38,
	}
	got, err := c.GetWeather("London,UK")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCelsiusCorrectlyConvertsKelvinToCelsius(t *testing.T) {
	t.Parallel()
	input := weather.Temperature(274.15)
	want := 1.0
	got := input.Celsius()
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
