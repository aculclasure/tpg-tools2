package howlong_test

import (
	"testing"
	"time"

	"github.com/aculclasure/howlong"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	testscript.RunMain(m, map[string]func() int{
		"howlong": howlong.Main,
	})
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestRunReportsCorrectElapsedTimeForCommand(t *testing.T) {
	t.Parallel()
	goal := 100 * time.Millisecond
	elapsed, err := howlong.Run("sleep", "0.1")
	if err != nil {
		t.Fatal(err)
	}
	delta := goal - elapsed
	if delta.Abs() > 300*time.Millisecond {
		t.Errorf("want %s, got %s (not close enough)", goal, elapsed)
	}
}
