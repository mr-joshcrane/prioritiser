package prioritiser

import (
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

func GetUserPreference[T any](a, b T, r io.Reader, w io.Writer) T {
	s := ""
	fmt.Fprintf(w, "Do you prefer %v over %v?\n", a, b)
	for {
		fmt.Fscan(r, &s)
		if s == "y" {
			return a
		}
		if s == "n" {
			return b
		}
		fmt.Fprintf(w, "Invalid response recieved %s include (y) or (n)\n", s)
	}
}

func RunCLI(items []string, r io.Reader, w io.Writer) {
	sort.Slice(items, func(i, j int) bool {
		pref := GetUserPreference(items[i], items[j], r, w)
		return pref == items[i]
	})
	fmt.Fprintln(w, items)
}

func RandomList() []string {
	list := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	return list
}
