package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mr-joshcrane/prioritiser"
)

func main() {
	addMode := flag.Bool("add", false, "Adds a new item to an already sorted list")
	flag.Parse()
	if !(len(os.Args) > 1) {
		fmt.Fprintf(os.Stderr, "Please supply file path")
		os.Exit(1)
	}

	input, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open '%s'\n", flag.Arg(0))
		os.Exit(1)
	}
	p := prioritiser.NewPrioritiser(prioritiser.WithInput(input), prioritiser.WithAddMode(*addMode))

	p.RunCLI()
}
