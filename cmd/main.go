package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mr-joshcrane/prioritiser"
)

func main() {
	addMode := flag.String("add", "", "Adds a new item to an already sorted list")
	flag.Parse()
	if !(len(os.Args) > 1)  {
		fmt.Fprintf(os.Stderr, "Please supply file path")
		os.Exit(1)
	}
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open %s", os.Args[1])
		os.Exit(1)
	}
	s := strings.Split(string(input), "\n")
	t := []string{}
	for _, v := range s {
		if v != "" {
			t = append(t, v)
		}
	}
	priorities := prioritiser.WithPriorities(t)
	if *addMode != "" {
		priorities = prioritiser.WithPriorPriorities(t)

	}

	p := prioritiser.NewPrioritiser(priorities)
	p.RunCLI()
}
