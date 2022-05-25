package prioritiser

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/google/go-cmp/cmp"
)

type Prioritiser[T Prioritisables] struct {
	r               io.Reader
	w               io.Writer
	priorities      []T
	priorPriorities []T
}

type Prioritisables interface {
	~string
}

type PrioritiserOption[T Prioritisables] func(*Prioritiser[T]) *Prioritiser[T]

func WithReader[T Prioritisables](r io.Reader) PrioritiserOption[T] {
	return func(p *Prioritiser[T]) *Prioritiser[T] {
		p.r = r
		return p
	}
}

func WithWriter[T Prioritisables](w io.Writer) PrioritiserOption[T] {
	return func(p *Prioritiser[T]) *Prioritiser[T] {
		p.w = w
		return p
	}
}

func WithPriorities[T Prioritisables](priorities []T) PrioritiserOption[T] {
	return func(p *Prioritiser[T]) *Prioritiser[T] {
		p.priorities = priorities
		return p
	}
}

func WithPriorPriorities[T Prioritisables](priorPriorities []T) PrioritiserOption[T] {
	return func(p *Prioritiser[T]) *Prioritiser[T] {
		p.priorPriorities = priorPriorities
		return p
	}
}

func NewPrioritiser[T Prioritisables](opts ...PrioritiserOption[T]) *Prioritiser[T] {
	p := Prioritiser[T]{
		r: os.Stdin,
		w: os.Stdout,
	}
	for _, opt := range opts {
		opt(&p)
	}
	return &p
}

func (p *Prioritiser[T]) GetUserPreference(a, b T) T {
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

func (p *Prioritiser[T]) Sort() []T {
	sort.Slice(p.priorities, func(i, j int) bool {
		pref := p.GetUserPreference(p.priorities[i], p.priorities[j])
		return pref == p.priorities[i]
	})
	return p.priorities
}

func (p *Prioritiser[T]) MergeOne(item T, l []T) []T {
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

func (p *Prioritiser[T]) MergeLists() []T {
	for i := 0; i < len(p.priorities); i++ {
		p.priorPriorities = p.MergeOne(p.priorities[i], p.priorPriorities)
	}
	return p.priorPriorities
}

func (p *Prioritiser[string]) RunCLI() []string {
	if p.priorities == nil {
		p.priorities = func() []string {
			var items []string
			for {
				fmt.Fprintf(p.w, "Please add your new item.\n")
				fmt.Fprintf(p.w, "If you have no more items, press enter.\n")
				scanner := bufio.NewScanner(p.r)
				scanner.Scan()
				line := scanner.Text()

				if cmp.Equal(line, "") {
					break
				}
				items = append(items, line)
			}
			return items
		}()
	}
	p.Sort()
	if p.priorPriorities != nil {
		p.MergeLists()
	}
	fmt.Fprintln(p.w, "Sorted Priorities:")
	for _, v := range p.priorPriorities {
		fmt.Fprintln(p.w, v)
	}
	return p.priorPriorities
}

func RandomList() []string {
	list := []string{"1", "2", "3", "4", "5"}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	return list
}
