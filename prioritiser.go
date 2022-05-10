package prioritiser

import (
	"fmt"
	"io"
	"os"
)

func Sort(input []string, r io.Reader) []string {
	return insertionSort(input, r)
}

func insertionSort(arr []string, r io.Reader) []string {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && UserComparison(arr[j-1], arr[j], r); j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func UserComparison(a, b string, r io.Reader) bool {
	s := ""
	fmt.Printf("Is %s > %s?\n", a, b)
	fmt.Fscan(r, &s)
	if s == "y" {
		return true
	}
	if s == "n" {
		return false
	}
	fmt.Printf("Invalid response recieved %s include (y) or (n)\n", s)
	return UserComparison(a, b, r)
}

func RunCLI(input []string) {
	sorted := Sort(input, os.Stdin)
	fmt.Println(sorted)
}
