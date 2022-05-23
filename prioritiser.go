package prioritiser

import (
	"fmt"
	"golang.org/x/exp/slices"
	"io"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Prioritiser[T comparable] struct {
	r               io.Reader
	w               io.Writer
	priorities      []T
	priorPriorities []T
}

type PrioritiserOption[T comparable] func(*Prioritiser[T]) *Prioritiser[T]

func WithReader[T comparable](r io.Reader) PrioritiserOption[T] {
	return func(p *Prioritiser[T]) *Prioritiser[T] {
		p.r = r
		return p
	}
}

func WithWriter[T comparable](w io.Writer) PrioritiserOption[T] {
	return func(p *Prioritiser[T]) *Prioritiser[T] {
		p.w = w
		return p
	}
}

func WithPriorities[T comparable](priorities []T) PrioritiserOption[T] {
	return func(p *Prioritiser[T]) *Prioritiser[T] {
		p.priorities = priorities
		return p
	}
}

func WithPriorPriorities[T comparable](priorPriorities []T) PrioritiserOption[T] {
	return func(p *Prioritiser[T]) *Prioritiser[T] {
		p.priorPriorities = priorPriorities
		return p
	}
}

func NewPrioritiser[T comparable](opts ...PrioritiserOption[T]) *Prioritiser[T] {
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
	fmt.Fprintf(p.w, "Do you prefer %v over %v?\n", a, b)
	for {
		fmt.Fscan(p.r, &s)
		if s == "y" {
			return a
		}
		if s == "n" {
			return b
		}
		fmt.Fprintf(p.w, "Invalid response recieved %s include (y) or (n)\n", s)
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
		if slices.Contains(l, items[i]) && slices.Contains(l, items[j]) {
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

func (p *Prioritiser[T]) RunCLI() {
	result := p.Sort()
	if p.priorPriorities != nil {
		result = p.MergeLists()
	}
	fmt.Fprintln(p.w, result)
}

func RandomList() []string {
	list := []string{"1", "2", "3", "4", "5"}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	return list
}
