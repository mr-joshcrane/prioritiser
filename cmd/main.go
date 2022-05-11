package main

import (
	"os"

	"github.com/mr-joshcrane/prioritiser"
)

func main() {
	input := []string{"important task", "not important task"}
	prioritiser.RunCLI(input, os.Stdin, os.Stdout)
}
