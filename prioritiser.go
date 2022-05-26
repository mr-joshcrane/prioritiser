package prioritiser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

type Prioritiser struct {
	r                  io.Reader
	w                  io.Writer
	input              io.Reader
	addMode            bool
	unsortedPriorities []string
	sortedPriorities   []string
	lookupTable        map[string]int
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

func WithInput(input io.Reader) PrioritiserOption {
	return func(p *Prioritiser) *Prioritiser {
		p.input = input
		return p
	}
}

func WithAddMode(addMode bool) PrioritiserOption {
	return func(p *Prioritiser) *Prioritiser {
		p.addMode = addMode
		return p
	}
}

func WithUnsortedPriorities(priorities []string) PrioritiserOption {
	return func(p *Prioritiser) *Prioritiser {
		p.unsortedPriorities = priorities
		return p
	}
}

func WithSortedPriorities(priorPriorities []string) PrioritiserOption {
	return func(p *Prioritiser) *Prioritiser {
		p.sortedPriorities = priorPriorities
		return p
	}
}

func NewPrioritiser(opts ...PrioritiserOption) *Prioritiser {
	p := Prioritiser{
		r:           os.Stdin,
		w:           os.Stdout,
		lookupTable: map[string]int{},
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

func (p *Prioritiser) GetUserPreferenceBS(a, b string) int {
	if val, ok := p.lookupTable[a+b]; ok {
		return val
	}
	if val, ok := p.lookupTable[b+a]; ok {
		return val
	}
	s := ""
	fmt.Fprintf(p.w, "\n1. %v\nOR\n2. %v?\n", a, b)
	for {
		fmt.Fscan(p.r, &s)
		fmt.Fprintf(p.w, "Selected %s\n", s)
		if s == "1" {
			p.lookupTable[a+b] = 1
			return p.lookupTable[a+b]
		}
		if s == "2" {
			p.lookupTable[b+a] = -1
			return p.lookupTable[b+a]
		}
		fmt.Fprintf(p.w, "Invalid response recieved %s include (1) or (2)\n", s)
	}
}

func (p *Prioritiser) Sort() []string {
	sort.Slice(p.unsortedPriorities, func(i, j int) bool {
		pref := p.GetUserPreference(p.unsortedPriorities[i], p.unsortedPriorities[j])
		return pref == p.unsortedPriorities[i]
	})
	return p.unsortedPriorities
}

func (p *Prioritiser) MergeOne(item string, l []string) []string {
	i, _ := slices.BinarySearchFunc(l, item, p.GetUserPreferenceBS)
	return slices.Insert(l, i, item)
}

func (p *Prioritiser) MergeLists() []string {
	for i := 0; i < len(p.unsortedPriorities); i++ {
		p.sortedPriorities = reverse(p.MergeOne(p.unsortedPriorities[i], reverse(p.sortedPriorities)))
	}
	return p.sortedPriorities
}

func (p Prioritiser) GetUserPriorities() []string {
	var items []string
	scanner := bufio.NewScanner(p.r)
	for {
		line := ""
		fmt.Fprintf(p.w, "Please add your new item.\n")
		fmt.Fprintf(p.w, "To exit, type Q and enter.\n")
		scanner.Scan()
		line = scanner.Text()
		if line == "q" || line == "Q" {
			break
		}
		isPreviousPriority := slices.Contains(p.sortedPriorities, line)
		isDuplicateEntry := slices.Contains(items, line)
		if line != "" && !isPreviousPriority && !isDuplicateEntry {
			items = append(items, line)
		}
	}
	return items
}

func ManagePriorities(p *Prioritiser) []string {
	if p.unsortedPriorities == nil {
		p.unsortedPriorities = p.GetUserPriorities()
	}
	result := p.Sort()
	if p.sortedPriorities != nil {
		result = p.MergeLists()
	}
	return result
}

func OutputPriorities(w io.Writer, priorities []string) {
	fmt.Fprintln(w, "Sorted Priorities:")
	for _, v := range priorities {
		fmt.Fprintln(w, v)
	}
}

func ValidateInput(input io.Reader) []string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(input)
	s := strings.Split(buf.String(), "\n")
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == "" {
			s = slices.Delete(s, i, i+1)
		}
	}
	return s
}

func (p *Prioritiser) RunCLI() {
	s := ValidateInput(p.input)
	if p.addMode {
		p.sortedPriorities = s
	} else {
		p.unsortedPriorities = s
	}
	priorities := ManagePriorities(p)
	OutputPriorities(p.w, priorities)
}

func reverse(input []string) []string {
	var output []string
	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}
	return output
}
