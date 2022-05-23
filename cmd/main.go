package main

import (
	"github.com/mr-joshcrane/prioritiser"
)

func main() {
	input := prioritiser.RandomList()
	priors := []string{"12", "11", "10"}
	
	priorities := prioritiser.WithPriorities(input)
	priorPriorities := prioritiser.WithPriorPriorities(priors)

	p := prioritiser.NewPrioritiser(priorities, priorPriorities)
	p.RunCLI()
}
