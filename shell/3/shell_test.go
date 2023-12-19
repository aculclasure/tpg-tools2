package shell_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/aculclasure/shell"
	"github.com/google/go-cmp/cmp"
)

func TestCmdFromString_CreatesExpectedCmd(t *testing.T) {
	t.Parallel()
	input := "/bin/ls -l"
	want := []string{"/bin/ls", "-l"}
	cmd, err := shell.CmdFromString(input)
	if err != nil {
		t.Fatal(err)
	}
	got := cmd.Args
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCmdFromString_ErrorsOnEmptyInput(t *testing.T) {
	t.Parallel()
	_, err := shell.CmdFromString("")
	if err == nil {
		t.Fatal("want error on empty input but got nil")
	}
}

func TestNewSession_CreatesExpectedSession(t *testing.T) {
	t.Parallel()
	want := shell.Session{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)
	if want != got {
		t.Errorf("want %#v, got %#v", want, got)
	}
}

func TestRunProducesExpectedOutput(t *testing.T) {
	t.Parallel()
	in := strings.NewReader("echo hello\n\n")
	out := new(bytes.Buffer)
	session := shell.NewSession(in, out, io.Discard)
	session.DryRun = true
	session.Run()
	want := "> echo hello\n> > \nBe seeing you!\n"
	got := out.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
