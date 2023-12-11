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
		"count": count.Main,
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

func TestCountWords_CountsWordsInInput(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("one word\nfour words\nsix words")
	counter, err := count.NewCounter(
		count.WithInput(input),
	)
	if err != nil {
		t.Fatal(err)
	}
	got := counter.CountWords()
	want := 6
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCountBytes_CountsBytesInInput(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("1")
	counter, err := count.NewCounter(
		count.WithInput(input),
	)
	if err != nil {
		t.Fatal(err)
	}
	got := counter.CountBytes()
	want := 1
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInput_ReturnsErrorForNilInputReader(t *testing.T) {
	t.Parallel()
	_, err := count.NewCounter(
		count.WithInput(nil),
	)
	if err == nil {
		t.Error("wanted an error but did not get one")
	}
}

func TestWithInputFromArgs_IgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("1\n2\n3")
	counter, err := count.NewCounter(
		count.WithInput(input),
		count.WithInputFromArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := counter.CountLines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs_SetsInputToPath(t *testing.T) {
	t.Parallel()
	counter, err := count.NewCounter(
		count.WithInputFromArgs([]string{"testdata/three_lines.txt"}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := counter.CountLines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestCountLines_WithMultiplePathsReturnsTotalNumberOfLines(t *testing.T) {
	t.Parallel()
	counter, err := count.NewCounter(
		count.WithInputFromArgs([]string{
			"testdata/three_lines.txt",
			"testdata/four_lines.txt",
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 7
	got := counter.CountLines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
