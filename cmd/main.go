package main

import (
	"github.com/mr-joshcrane/prioritiser"
)

func main() {
	input := prioritiser.RandomList()
	
	priorities := prioritiser.WithPriorities(input)
	p := prioritiser.NewPrioritiser(priorities)

	// An example of how you might take an already sorted list to integrate new items into
	// priors := []string{"12", "11", "10"}
	// priorPriorities := prioritiser.WithPriorPriorities(priors)
	// p := prioritiser.NewPrioritiser(priorities, priorPriorities)

	p.RunCLI()
}
