package howlong

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const usage = `Usage: howlong COMMAND [ARGS...]

Runs COMMAND with ARGS and reports the elapsed wall-clock time.`

func Run(program string, args ...string) (time.Duration, error) {
	cmd := exec.Command(program, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	start := time.Now()
	err := cmd.Run()
	elapsed := time.Since(start)
	if err != nil {
		return 0, err
	}
	return elapsed, nil
}

func Main() int {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		return 0
	}
	elapsed, err := Run(os.Args[1], os.Args[2:]...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Printf("(time: %s)\n", elapsed.Round(time.Millisecond))
	return 0
}
