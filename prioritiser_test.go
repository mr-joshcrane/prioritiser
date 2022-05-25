package prioritiser_test

import (
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mr-joshcrane/prioritiser"
)

func TestGetUserCanAddNewItemsToPreviouslySortedPriorities(t *testing.T) {
	t.Parallel()
	reader := strings.NewReader("2\n1\n1\n2\n1\n2")
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(io.Discard)
	pOption := prioritiser.WithPriorities([]string{"2", "4"})
	ppOption := prioritiser.WithPriorPriorities([]string{"1", "3", "5"})
	p := prioritiser.NewPrioritiser(readerOption, writerOption, ppOption, pOption)

	got := prioritiser.ManagePriorities(p)
	want := []string{"1", "2", "3", "4", "5"}
	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

// func TestGetUserPreferenceGivenAAndBReturnsAIfUserEntersY(t *testing.T) {
// 	t.Parallel()
// 	reader := strings.NewReader("y")
// 	readerOption := prioritiser.WithReader[string](reader)
// 	writerOption := prioritiser.WithWriter[string](io.Discard)
// 	p := prioritiser.NewPrioritiser(readerOption, writerOption)

// 	got := p.GetUserPreference("preferred", "non-preferred")
// 	want := "preferred"
// 	if want != got {
// 		t.Fatalf("wanted %v, got %v", want, got)
// 	}

// func TestGetUserPreferenceGivenInvalidInput(t *testing.T) {
// 	t.Parallel()
// 	reader := strings.NewReader("ziggy\nwalrus\nnn\ny")
// 	readerOption := prioritiser.WithReader[string](reader)
// 	writerOption := prioritiser.WithWriter[string](io.Discard)
// 	p := prioritiser.NewPrioritiser(readerOption, writerOption)

// 	got := p.GetUserPreference("preferred", "non-preferred")
// 	want := "preferred"
// 	if want != got {
// 		t.Fatalf("wanted %v, got %v", want, got)
// 	}
// }

// func TestRunCLIReturnsItemsInDescendingOrderOfImportance(t *testing.T) {
// 	t.Parallel()
// 	reader := strings.NewReader("n\ny\nn")
// 	buf := &bytes.Buffer{}
// 	input := []string{"high", "low", "medium"}

// 	readerOption := prioritiser.WithReader[string](reader)
// 	writerOption := prioritiser.WithWriter[string](buf)
// 	pOption := prioritiser.WithPriorities(input)
// 	p := prioritiser.NewPrioritiser(readerOption, writerOption, pOption)

// 	p.RunCLI()

// 	want := "Do you prefer low over high?\nDo you prefer medium over low?\nDo you prefer medium over high?\n[high medium low]\n"
// 	got := buf.String()

// 	if want != got {
// 		t.Fatalf("wanted %v, got %v", want, got)
// 	}
// }

// func TestCanMergePreviouslySortedList(t *testing.T) {
// 	t.Parallel()
// 	reader := strings.NewReader("y\ny\ny\nn\ny\nn")
// 	priorPriorities := []string{"high", "medium", "low"}
// 	priorities := []string{"highest", "lowest"}

// 	readerOption := prioritiser.WithReader[string](reader)
// 	writerOption := prioritiser.WithWriter[string](io.Discard)
// 	ppOption := prioritiser.WithPriorPriorities(priorPriorities)
// 	pOption := prioritiser.WithPriorities(priorities)

// 	p := prioritiser.NewPrioritiser(readerOption, writerOption, pOption, ppOption)

// 	want := []string{"highest", "high", "medium", "low", "lowest"}
// 	got := p.MergeLists()
// 	if !cmp.Equal(want, got) {
// 		t.Fatalf("wanted %v, got %v", want, got)
// 	}
// }
