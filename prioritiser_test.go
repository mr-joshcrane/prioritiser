package prioritiser_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestGetUserPreferenceGivenInvalidInput(t *testing.T) {
	t.Parallel()
	r := strings.NewReader("ziggy\nwalrus\nnn\ny")
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
	r := strings.NewReader("n\ny\nn")

	prioritiser.RunCLI(input, nil, r, buf)
	got := buf.String()

	if want != got {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCanMergePreviouslySortedList(t *testing.T) {
	t.Parallel()
	prevSorted := []string{"high", "medium", "low"}
	newItems := []string{"highest", "lowest"}
	r := strings.NewReader("y\ny\ny\nn\ny\nn")
	want := []string{"highest", "high", "medium", "low", "lowest"}
	got := prioritiser.MergeLists(newItems, prevSorted, r, io.Discard)
	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
