package count

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input io.Reader
	files []io.Reader
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
		c.files = make([]io.Reader, len(args))
		for i, path := range args {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			c.files[i] = f
		}
		c.input = io.MultiReader(c.files...)
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
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return lines
}

func (c *counter) CountWords() int {
	scn := bufio.NewScanner(c.input)
	scn.Split(bufio.ScanWords)
	words := 0
	for scn.Scan() {
		scn.Text()
		words++
	}
	for _, f := range c.files {
		f.(io.Closer).Close()
	}
	return words
}

func Main() int {
	lineMode := flag.Bool("lines", false, "Count lines, not words")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-lines] [files...]\n", os.Args[0])
		fmt.Println("Count words (or lines) from stdin (or files).")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	flag.Parse()
	c, err := NewCounter(WithInputFromArgs(flag.Args()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if *lineMode {
		fmt.Println(c.CountLines())
	} else {
		fmt.Println(c.CountWords())
	}
	return 0
}
