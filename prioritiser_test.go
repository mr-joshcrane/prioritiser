package prioritiser_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/mr-joshcrane/prioritiser"
)

func TestGetUserPreferenceGivenAAndBReturnsAIfUserEntersY(t *testing.T) {
	t.Parallel()
	r := strings.NewReader("y")
	got := prioritiser.GetUserPreference("preferred", "non-preferred", r, io.Discard)
	want := "preferred"
	if want != got {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestRunCLIReturnsItemsInDescendingOrderOfImportance(t *testing.T) {
	t.Parallel()
	input := []string{"high", "low", "medium"}
	buf := &bytes.Buffer{}

	want := "Do you prefer low over high?\nDo you prefer medium over low?\nDo you prefer medium over high?\n[high medium low]\n"
	r := strings.NewReader("n\ny\nn\n")
	
	prioritiser.RunCLI(input, r, buf)
	got := buf.String()
	if want != got {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCanInsertNewItemsIntoPreviouslySortedList(t *testing.T) {
	t.Parallel()
	prevSorted := []string{"high", "low", "medium"}
	newItems := []string{"highest", "lowest"}
	want := []string{"highest", "high", "medium",  "low", "lowest"}
	got := prioritiser.Sort(prevSorted, newItems)
	if want != got {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}


