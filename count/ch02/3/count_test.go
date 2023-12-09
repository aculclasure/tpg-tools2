package count_test

import (
	"bytes"
	"testing"

	"github.com/aculclasure/count"
)

func TestCountLines_CountsLinesInInput(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("1\n2\n3")
	counter, err := count.NewCounter(
		count.WithInput(input),
	)
	if err != nil {
		t.Fatal(err)
	}
	got := counter.CountLines()
	want := 3
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInput_ReturnsErrorForNilInput(t *testing.T) {
	t.Parallel()
	_, err := count.NewCounter(
		count.WithInput(nil),
	)
	if err == nil {
		t.Error("wanted an error but did not get one")
	}
}
