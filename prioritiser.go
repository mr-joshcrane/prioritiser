package prioritiser

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

type Prioritiser struct {
	r               io.Reader
	w               io.Writer
	priorities      []string
	priorPriorities []string
}
type PrioritiserOption func(*Prioritiser) *Prioritiser

func WithReader(r io.Reader) PrioritiserOption {
	return func(p *Prioritiser) *Prioritiser {
		p.r = r
		return p
	}
}

func WithWriter(w io.Writer) PrioritiserOption {
	return func(p *Prioritiser) *Prioritiser {
		p.w = w
		return p
	}
}

func WithPriorities(priorities []string) PrioritiserOption {
	return func(p *Prioritiser) *Prioritiser {
		p.priorities = priorities
		return p
	}
}

func WithPriorPriorities(priorPriorities []string) PrioritiserOption {
	return func(p *Prioritiser) *Prioritiser {
		p.priorPriorities = priorPriorities
		return p
	}
}

func NewPrioritiser(opts ...PrioritiserOption) *Prioritiser {
	p := Prioritiser{
		r: os.Stdin,
		w: os.Stdout,
	}
	for _, opt := range opts {
		opt(&p)
	}
	return &p
}

func (p *Prioritiser) GetUserPreference(a, b string) string {
	s := ""
	fmt.Fprintf(p.w, "\n1. %v\nOR\n2. %v?\n", a, b)
	for {
		fmt.Fscan(p.r, &s)
		fmt.Fprintf(p.w, "Selected %s\n", s)
		if s == "1" {
			return a
		}
		if s == "2" {
			return b
		}
		fmt.Fprintf(p.w, "Invalid response recieved %s include (1) or (2)\n", s)
	}
}

func (p *Prioritiser) Sort() []string {
	sort.Slice(p.priorities, func(i, j int) bool {
		pref := p.GetUserPreference(p.priorities[i], p.priorities[j])
		return pref == p.priorities[i]
	})
	return p.priorities
}

func (p *Prioritiser) MergeOne(item string, l []string) []string {
	items := append(l, item)
	sort.Slice(items, func(i, j int) bool {
		if (item != items[i]) && (item != items[j]) {
			return i < j
		}
		pref := p.GetUserPreference(items[i], items[j])
		return pref == items[i]
	})
	return items
}

func (p *Prioritiser) MergeLists() []string {
	for i := 0; i < len(p.priorities); i++ {
		p.priorPriorities = p.MergeOne(p.priorities[i], p.priorPriorities)
	}
	return p.priorPriorities
}

func (p *Prioritiser) GetUserPriorities() []string {
	var items []string
	for {
		line := ""
		fmt.Fprintf(p.w, "Please add your new item.\n")
		fmt.Fprintf(p.w, "To exit, type Q and enter.\n")
		fmt.Fscan(p.r, &line)
		if line == "q" || line == "Q" {
			break
		}
		if line != "" {
			items = append(items, line)
		}
	}
	return items
}

func RunCLI() {
	addMode := flag.Bool("add", false, "Adds a new item to an already sorted list")
	flag.Parse()

	if !(len(os.Args) > 1) {
		fmt.Fprintf(os.Stderr, "Please supply file path")
		os.Exit(1)
	}
	input, err := os.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open '%s'\n", flag.Arg(0))
		os.Exit(1)
	}
	s := strings.Split(string(input), "\n")
	t := []string{}
	for _, v := range s {
		if v != "" {
			t = append(t, v)
		}
	}
	priorities := WithPriorities(t)
	if *addMode {
		priorities = WithPriorPriorities(t)
	}
	p := NewPrioritiser(priorities)

	if p.priorities == nil {
		p.priorities = p.GetUserPriorities()
	}
	p.Sort()
	if p.priorPriorities != nil {
		p.MergeLists()
	}
	fmt.Fprintln(p.w, "Sorted Priorities:")
	for _, v := range p.priorPriorities {
		fmt.Fprintln(p.w, v)
	}
}

func RandomList() []string {
	list := []string{"1", "2", "3", "4", "5"}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	return list
}
