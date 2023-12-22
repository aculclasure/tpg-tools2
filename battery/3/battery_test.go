package battery_test

import (
	"os"
	"testing"

	"github.com/aculclasure/battery"
	"github.com/google/go-cmp/cmp"
)

func TestParseAcpiOutput_GetsChargePercent(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/acpi.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := battery.Status{
		ChargePercent: 100,
	}
	got, err := battery.ParseAcpiOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestToJSON_GivesExpectedJSON(t *testing.T) {
	t.Parallel()
	batt := battery.Battery{
		Name:          "Battery0",
		ChargePercent: 100,
	}
	wantBytes, err := os.ReadFile("testdata/battery.json")
	if err != nil {
		t.Fatal(err)
	}
	want := string(wantBytes)
	got := batt.ToJSON()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
