package pipeline_test

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/aculclasure/pipeline"
	"github.com/google/go-cmp/cmp"
)

func TestStdoutPrintsMessageToOutput(t *testing.T) {
	t.Parallel()
	want := "Hello, world!\n"
	p := pipeline.FromString(want)
	out := new(bytes.Buffer)
	p.Output = out
	p.Stdout()
	if p.Err != nil {
		t.Fatal(p.Err)
	}
	got := out.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestStdoutPrintsNothingOnError(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello, world!\n")
	out := new(bytes.Buffer)
	p.Output = out
	p.Err = errors.New("oh no!")
	p.Stdout()
	got := out.String()
	want := ""
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFromFile_SetsErrorGivenNonexistentFile(t *testing.T) {
	t.Parallel()
	p := pipeline.FromFile("bogus.txt")
	if p.Err == nil {
		t.Fatal("wanted an error but did not get one for nonexistent file")
	}
}

func TestFromFile_ReadsAllDataFromFile(t *testing.T) {
	t.Parallel()
	p := pipeline.FromFile("testdata/hello.txt")
	if p.Err != nil {
		t.Fatal(p.Err)
	}
	want, err := os.ReadFile("testdata/hello.txt")
	if err != nil {
		t.Fatal(err)
	}
	got, err := io.ReadAll(p.Input)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}
