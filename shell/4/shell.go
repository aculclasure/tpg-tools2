package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type Session struct {
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
	Transcript io.Writer
	DryRun     bool
}

func NewSession(stdin io.Reader, stdout, stderr io.Writer) *Session {
	return &Session{
		Stdin:      stdin,
		Stdout:     stdout,
		Stderr:     stderr,
		Transcript: io.Discard,
	}
}

func (s *Session) Run() {
	stdout := io.MultiWriter(s.Stdout, s.Transcript)
	stderr := io.MultiWriter(s.Stderr, s.Transcript)
	fmt.Fprintf(stdout, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		fmt.Fprintln(s.Transcript, line)
		cmd, err := CmdFromString(line)
		if err != nil {
			fmt.Fprintf(stdout, "> ")
			continue
		}
		if s.DryRun {
			fmt.Fprintf(stdout, "%s\n> ", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(stderr, "error:", err)
		}
		fmt.Fprintf(stdout, "%s> ", output)
	}
	fmt.Fprintln(stdout, "\nBe seeing you!")
}

func CmdFromString(input string) (*exec.Cmd, error) {
	args := strings.Fields(input)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}
