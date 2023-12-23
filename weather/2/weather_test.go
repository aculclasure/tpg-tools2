package weather_test

import (
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
		Summary: "Clouds",
	}
	got, err := weather.ParseData(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		cmp.Diff(want, got)
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
