package count

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input io.Reader
}

type option func(*counter)

func WithInput(input io.Reader) option {
	return func(c *counter) {
		c.input = input
	}
}

func NewCounter(opts ...option) *counter {
	c := &counter{
		input: os.Stdin,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *counter) CountLines() int {
	scn := bufio.NewScanner(c.input)
	lines := 0
	for scn.Scan() {
		scn.Text()
		lines++
	}
	return lines
}

func Main() int {
	fmt.Println(NewCounter().CountLines())
	return 0
}
