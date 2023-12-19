package pipeline

import (
	"io"
	"os"
	"strings"
)

type Pipeline struct {
	Input  io.Reader
	Output io.Writer
	Err    error
}

func FromString(input string) *Pipeline {
	return &Pipeline{
		Input:  strings.NewReader(input),
		Output: os.Stdout,
	}
}

func (p *Pipeline) Stdout() {
	if p.Err != nil {
		return
	}
	io.Copy(p.Output, p.Input)
}
