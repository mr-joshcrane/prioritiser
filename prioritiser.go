package prioritiser

import (
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"

	"golang.org/x/exp/slices"
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

func Sort(items []string, r io.Reader, w io.Writer) []string {
	sort.Slice(items, func(i, j int) bool {
		pref := GetUserPreference(items[i], items[j], r, w)
		return pref == items[i]
	})
	return items
}

func SortSingle(single string, l []string, r io.Reader, w io.Writer) []string {
	items := append(l, single)
	sort.Slice(items, func(i, j int) bool {
		if slices.Contains(l, items[i]) && slices.Contains(l, items[j]) {
			return i < j
		}
		pref := GetUserPreference(items[i], items[j], r, w)
		return pref == items[i]
	})
	return items
}

func RunCLI(items []string, oldItems []string, r io.Reader, w io.Writer) {
	result := Sort(items, r, w)
	if oldItems != nil {
		result = MergeLists(result, oldItems, r, w)
	}
	fmt.Fprintln(w, result)
}

func RandomList() []string {
	list := []string{"1", "2", "3"}
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	return list
}

func MergeLists(newItems, previousItems []string, r io.Reader, w io.Writer) []string {
	for i := 0; i < len(newItems); i++ {
		previousItems = SortSingle(newItems[i], previousItems, r, w)
	}
	return previousItems
}
