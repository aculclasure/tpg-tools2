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
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
	DryRun bool
}

func NewSession(stdin io.Reader, stdout, stderr io.Writer) *Session {
	return &Session{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
	}
}

func (s *Session) Run() {
	fmt.Fprintf(s.Stdout, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := CmdFromString(line)
		if err != nil {
			fmt.Fprintf(s.Stdout, "> ")
			continue
		}
		if s.DryRun {
			fmt.Fprintf(s.Stdout, "%s\n> ", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(s.Stderr, "error:", err)
		}
		fmt.Fprintf(s.Stdout, "%s> ", output)
	}
	fmt.Fprintln(s.Stdout, "\nBe seeing you!")
}

func CmdFromString(input string) (*exec.Cmd, error) {
	args := strings.Fields(input)
	if len(args) < 1 {
		return nil, errors.New("empty input")
	}
	return exec.Command(args[0], args[1:]...), nil
}
