package main

import (
	"os"

	"github.com/mr-joshcrane/prioritiser"
)

func main() {
	input := prioritiser.RandomList()
	prioritiser.RunCLI(input, nil, os.Stdin, os.Stdout)
}
