package main

import "github.com/aculclasure/pipeline"

func main() {
	pipeline.FromString("Hello world!\n").Stdout()
}
