package greet_test

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/aculclasure/greet"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"greet": greet.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestGreetUser_PromptsUserForANameAndRendersMessage(t *testing.T) {
	t.Parallel()
	input := strings.NewReader("Andrew")
	output := new(bytes.Buffer)
	greet.GreetUser(input, output)
	want := "What is your name?\nHello Andrew!\n"
	got := output.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestGreetUser_RendersHelloYouOnReadError(t *testing.T) {
	t.Parallel()
	input := iotest.ErrReader(errors.New("read error"))
	output := new(bytes.Buffer)
	greet.GreetUser(input, output)
	want := "What is your name?\nHello you!\n"
	got := output.String()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
