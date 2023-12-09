package count

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input io.Reader
}

type option func(*counter) error

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input source")
		}
		c.input = input
		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) < 1 {
			return nil
		}
		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		c.input = f
		return nil
	}
}

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		input: os.Stdin,
	}
	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *counter) CountLines() int {
	scn := bufio.NewScanner(c.input)
	lines := 0
	for scn.Scan() {
		scn.Text()
		lines++
	}
	v, ok := c.input.(io.Closer)
	if ok {
		v.Close()
	}
	return lines
}

func Main() int {
	c, err := NewCounter(WithInputFromArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(c.CountLines())
	return 0
}
