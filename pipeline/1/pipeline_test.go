package pipeline_test

import (
	"bytes"
	"errors"
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
