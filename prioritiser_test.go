package prioritiser_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mr-joshcrane/prioritiser"
)

func TestGetUserCanAddNewItemsToPreviouslySortedPriorities(t *testing.T) {
	t.Parallel()
	reader := strings.NewReader("2\n2\n1\n2\n2\n1\n1")
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(io.Discard)
	pOption := prioritiser.WithUnsortedPriorities([]string{"2", "4"})
	ppOption := prioritiser.WithSortedPriorities([]string{"1", "3", "5", "6", "7", "8"})
	p := prioritiser.NewPrioritiser(readerOption, writerOption, ppOption, pOption)

	got := prioritiser.ManagePriorities(p)
	want := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
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

func TestUserCantEnterDuplicatePriorities(t *testing.T) {
	t.Parallel()
	reader := strings.NewReader("great book\ngreat book\ngreat book\naverage book\nterrible book\nQ\n")
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(io.Discard)
	p := prioritiser.NewPrioritiser(readerOption, writerOption)

	got := p.GetUserPriorities()
	want := []string{"great book", "average book", "terrible book"}
	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestUserCantEnterPreviouslySortedPriorities(t *testing.T) {
	t.Parallel()
	reader := strings.NewReader("great book\naverage book\nterrible book\nQ\n")
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(io.Discard)
	ppOption := prioritiser.WithSortedPriorities([]string{"average book"})

	p := prioritiser.NewPrioritiser(readerOption, writerOption, ppOption)

	got := p.GetUserPriorities()
	want := []string{"great book", "terrible book"}
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
	buf := bytes.Buffer{}
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(&buf)
	inputOption := prioritiser.WithInput(input)
	addMode := prioritiser.WithAddMode(true)
	p := prioritiser.NewPrioritiser(readerOption, writerOption, inputOption, addMode)

	want := []string{"great book", "average book", "terrible book"}
	p.RunCLI()
	output := buf.String()
	got := strings.Split(output, "\n")
	got = got[len(got)-4 : len(got)-1]

	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCLIInNormalMode(t *testing.T) {
	t.Parallel()
	input := bytes.NewReader([]byte("great book\n\n\n\naverage book\nterrible book\n\n\n"))
	reader := strings.NewReader("2\n2\n2")
	buf := bytes.Buffer{}
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(&buf)
	inputOption := prioritiser.WithInput(input)
	p := prioritiser.NewPrioritiser(readerOption, writerOption, inputOption)

	want := []string{"great book", "average book", "terrible book"}
	p.RunCLI()
	output := buf.String()
	got := strings.Split(output, "\n")
	got = got[len(got)-4 : len(got)-1]
	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCLIInSaveMode(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/" + t.Name()
	input := bytes.NewReader([]byte("great book\n\n\n\naverage book\nterrible book\n\n\n"))
	reader := strings.NewReader("2\n2\n2")
	buf := bytes.Buffer{}
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(&buf)
	inputOption := prioritiser.WithInput(input)
	saveMode := prioritiser.WithSaveMode(true, path)
	p := prioritiser.NewPrioritiser(readerOption, writerOption, inputOption, saveMode)

	want := []string{"great book", "average book", "terrible book"}
	p.RunCLI()
	contents, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	got := strings.Split(string(contents), "\n")
	got = got[len(got)-4 : len(got)-1]
	if !cmp.Equal(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestCLIInSaveModeWhenFileCantBeWrittenTo(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/" + t.Name()
	input := bytes.NewReader([]byte("great book\n\n\n\naverage book\nterrible book\n\n\n"))
	reader := strings.NewReader("2\n2\n2")
	buf := bytes.Buffer{}
	readerOption := prioritiser.WithReader(reader)
	writerOption := prioritiser.WithWriter(&buf)
	inputOption := prioritiser.WithInput(input)
	saveMode := prioritiser.WithSaveMode(true, path)
	p := prioritiser.NewPrioritiser(readerOption, writerOption, inputOption, saveMode)

	_, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
 	err = os.Chmod(path, 0444)
	if err != nil {
		t.Fatal(err)
	}
	err = p.RunCLI()
	if err == nil {
		t.Fatalf("expected permission denied error, but got none")
	}
}