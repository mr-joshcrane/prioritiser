package prioritiser

import (
	"fmt"
	"io"
	"os"
	"sort"
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
		pref:= GetUserPreference(items[i], items[j], os.Stdin, os.Stdout)
		return pref == items[i] 
	})
	fmt.Println(items)
}
