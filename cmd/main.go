package main

import (
	"github.com/mr-joshcrane/prioritiser"
)

func main() {
	input := []string{"important task", "slightly important task", "most important task", "least important task"}
	prioritiser.RunCLI(input)
}
