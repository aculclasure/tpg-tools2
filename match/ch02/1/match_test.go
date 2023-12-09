package match_test

import (
	"bytes"
	"testing"

	"github.com/aculclasure/match"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.RunMain(m, map[string]func() int{
		"match": match.Main,
	})
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestMatch_MatchesLinesWithSubstringAndRendersMatchingLines(t *testing.T) {
	t.Parallel()
	input := bytes.NewBufferString("one two three\nfour five six\nseven one eight")
	output := new(bytes.Buffer)
	m, err := match.NewMatcher(
		match.WithInput(input),
		match.WithOutput(output),
		match.WithSearchTextFromArgs([]string{"one"}),
	)
	if err != nil {
		t.Fatal(err)
	}
	m.PrintMatchingLines()
	got := output.String()
	want := "one two three\nseven one eight\n"
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
