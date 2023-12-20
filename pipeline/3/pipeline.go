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

func New() *Pipeline {
	return &Pipeline{
		Output: os.Stdout,
	}
}

func FromString(input string) *Pipeline {
	p := New()
	p.Input = strings.NewReader(input)
	return p
}

func FromFile(path string) *Pipeline {
	f, err := os.Open(path)
	if err != nil {
		return &Pipeline{Err: err}
	}
	p := New()
	p.Input = f
	return p
}

func (p *Pipeline) Stdout() {
	if p.Err != nil {
		return
	}
	io.Copy(p.Output, p.Input)
}
