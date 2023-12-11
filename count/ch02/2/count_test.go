package count_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/aculclasure/count"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"countlines": count.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

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
