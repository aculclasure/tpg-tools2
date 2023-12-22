package weather_test

import (
	"os"
	"testing"

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
