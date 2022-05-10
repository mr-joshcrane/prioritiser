package prioritiser_test

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mr-joshcrane/prioritiser"
)

//tokeniser?
var Ordering = map[string]int{
	"most important task":     0,
	"important task":          1,
	"slightly important task": 2,
	"least important task":    3,
}

func TestSort(t *testing.T) {
	t.Parallel()
	var SortTests = []struct {
		description string
		input       []string
		responses   io.Reader
		want        []string
	}{
		{
			description: "Sorting a sorted list",
			input:       []string{"least important task", "slightly important task", "important task", "most important task"},
			responses:   strings.NewReader("n\nn\nn\nn\nn\nn\n"),
			want:        []string{"least important task", "slightly important task", "important task", "most important task"},
		},
		{
			description: "Sorting an unsorted list",
			input:       []string{"important task", "slightly important task", "most important task", "least important task"},
			responses:   strings.NewReader("y\nn\ny\ny\ny\n"),
			want:        []string{"least important task", "slightly important task", "important task", "most important task"},
		},
	}
	for _, tt := range SortTests {

		got := prioritiser.Sort(tt.input, tt.responses)
		if !cmp.Equal(tt.want, got) {
			t.Error(tt.description, cmp.Diff(tt.want, got))
		}
	}

}
