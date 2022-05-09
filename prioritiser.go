package prioritiser

import (
	"fmt"
)

func Sort(input []int) []int {
	return insertionSort(input)
}

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && UserComparison(arr[j-1], arr[j]); j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func UserComparison(a, b int) bool  {
	s := ""
	fmt.Printf("Is %d > %d?\n", a , b)
	fmt.Scanln(&s)
	if s == "y" {
		return true
	}
	if s == "n" {
		return false
	}
	fmt.Printf("Invalid response recieved %s include (y) or (n)\n", s)
	return UserComparison(a,b)
}


func RunCLI(input []int) {
	sorted := Sort(input)
	fmt.Println(sorted)
}
