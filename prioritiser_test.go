package prioritiser_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mr-joshcrane/prioritiser"
)

func TestSort(t *testing.T) {
	t.Parallel()
	var SortTests = []struct {
		description string
		input       []int
		want        []int
	}{
		{
			description: "Passing a sorted list",
			input:       []int{1, 2, 3, 4, 5},
			want:        []int{1, 2, 3, 4, 5},
		},
		{
			description: "Passing an unsorted list",
			input:       []int{2, 3, 1, 5, 4},
			want:        []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range SortTests {
		got := prioritiser.Sort(tt.input)
		if !cmp.Equal(tt.want, got) {
			t.Error(cmp.Diff(tt.want, got))
		}
	}

}
