package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mr-joshcrane/prioritiser"
)

func main() {
	addMode := flag.Bool("add", false, "Adds a new item to an already sorted list")
	saveMode := flag.Bool("i", false, "Edits the file in place")
	flag.Parse()
	if !(len(os.Args) > 1) {
		fmt.Fprintf(os.Stderr, "Please supply file path\nFor example: $ %s books.txt\n", os.Args[0])
		os.Exit(1)
	}
	path := flag.Arg(0)
	input, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open '%s'\n", path)
		os.Exit(1)
	}
	p := prioritiser.NewPrioritiser(prioritiser.WithInput(input), prioritiser.WithAddMode(*addMode), prioritiser.WithSaveMode(*saveMode, path))

	err = p.RunCLI()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to file '%s': %v\n", path, err)

	}
}
