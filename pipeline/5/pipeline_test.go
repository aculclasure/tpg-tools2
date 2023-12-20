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

func TestStringReturnsPipelineContents(t *testing.T) {
	t.Parallel()
	want := "Hello, world!"
	p := pipeline.FromString(want)
	got, err := p.String()
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}

}

func TestStringReturnsErrorWhenPipeErrorSet(t *testing.T) {
	t.Parallel()
	p := pipeline.New()
	p.Err = errors.New("oh no!")
	_, err := p.String()
	if err == nil {
		t.Fatal("want error from String when pipeline has error but got nil")
	}
}

func TestColumnSelectsColumn2of3(t *testing.T) {
	t.Parallel()
	input := "1 2 3\n1 2 3\n1 2 3"
	p := pipeline.FromString(input)
	want := "2\n2\n2\n"
	got, err := p.Column(2).String()
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestColumnProducesNothingWhenPipelineErrorSet(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("Hello, world!\n")
	p.Err = errors.New("1\n2\n3!")
	data, err := io.ReadAll(p.Column(1).Input)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column after error, but got %q", data)
	}
}

func TestColumnSetsPipelineErrorAndProducesNothingGivenInvalidArg(t *testing.T) {
	t.Parallel()
	p := pipeline.FromString("1 2 3\n1 2 3\n1 2 3")
	p.Column(-1)
	if p.Err == nil {
		t.Error("want an error on non-positive input but got nil")
	}
	data, err := io.ReadAll(p.Column(1).Input)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Errorf("want no output from Column with invalid col, but got %q", data)
	}
}
