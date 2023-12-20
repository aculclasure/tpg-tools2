package pipeline

import (
	"bufio"
	"fmt"
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

func (p *Pipeline) String() (string, error) {
	if p.Err != nil {
		return "", p.Err
	}
	data, err := io.ReadAll(p.Input)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (p *Pipeline) Column(col int) *Pipeline {
	if p.Err != nil {
		p.Input = strings.NewReader("")
		return p
	}
	if col < 1 {
		p.Err = fmt.Errorf("want col to have value of 1 or higher but got %d", col)
		return p
	}
	var result string
	buf := bufio.NewScanner(p.Input)
	for buf.Scan() {
		fields := strings.Fields(buf.Text())
		if len(fields) < col {
			continue
		}
		result += fields[col-1] + "\n"
	}
	return &Pipeline{
		Input: strings.NewReader(result),
	}
}
