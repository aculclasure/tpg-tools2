package greet

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GreetUser(input io.Reader, output io.Writer) {
	name := "you"
	fmt.Fprintln(output, "What is your name?")
	scn := bufio.NewScanner(input)
	if scn.Scan() {
		name = scn.Text()
	}
	fmt.Fprintf(output, "Hello %s!\n", name)
}

func Main() int {
	GreetUser(os.Stdin, os.Stdout)
	return 0
}
