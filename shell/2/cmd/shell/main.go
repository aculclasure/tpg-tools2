package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aculclasure/shell"
)

func main() {
	fmt.Print("> ")
	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		line := scn.Text()
		cmd, err := shell.CmdFromString(line)
		if err != nil {
			continue
		}
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf("%s", out)
		fmt.Print("\n> ")
	}
	fmt.Println("\nBe seeing you!")
}
