package prioritiser_test

import (
	"bytes"
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
	pOption := prioritiser.WithUnsortedPriorities([]string{"2", "4"})
	ppOption := prioritiser.WithSortedPriorities([]string{"1", "3", "5"})
	p := prioritiser.NewPrioritiser(readerOption, writerOption, ppOption, pOption)

	got := prioritiser.ManagePriorities(p)
	want := []string{"1", "2", "3", "4", "5"}
	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestGetUserPrioritiesCanGatherPrioritiesFromUserInput(t *testing.T) {
	t.Parallel()
	reader := strings.NewReader("great book\naverage book\nterrible book\nQ\n")
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(io.Discard)
	p := prioritiser.NewPrioritiser(readerOption, writerOption)

	got := p.GetUserPriorities()
	want := []string{"great book", "average book", "terrible book"}
	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestUserInputIsValidatedWithBlankItemsRemoved(t *testing.T) {
	t.Parallel()
	input := bytes.NewReader([]byte("great book\n\n\n\naverage book\nterrible book\n\n\n"))
	got := prioritiser.ValidateInput(input)
	want := []string{"great book", "average book", "terrible book"}
	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCLIInAddMode(t *testing.T) {
	t.Parallel()
	input := bytes.NewReader([]byte("great book\n\n\n\naverage book\nterrible book\n\n\n"))
	reader := strings.NewReader("Q\n")
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(io.Discard)
	inputOption := prioritiser.WithInput(input)
	addMode := prioritiser.WithAddMode(true)
	p := prioritiser.NewPrioritiser(readerOption, writerOption, inputOption, addMode)

	want := []string{"great book", "average book", "terrible book"}
	got := p.RunCLI()

	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCLIInNormalMode(t *testing.T) {
	t.Parallel()
	input := bytes.NewReader([]byte("great book\n\n\n\naverage book\nterrible book\n\n\n"))
	reader := strings.NewReader("2\n2\n")
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(io.Discard)
	inputOption := prioritiser.WithInput(input)
	p := prioritiser.NewPrioritiser(readerOption, writerOption, inputOption)

	want := []string{"great book", "average book", "terrible book"}
	got := p.RunCLI()

	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
