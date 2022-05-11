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
	got := prioritiser.GetUserPreference("preferred", "non-preferred" , r, io.Discard)
	want := "preferred"
	if want != got {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func TestRunCLIReturnsItemsInDescendingOrderOfImportance(t *testing.T) {
	input := []string{"high", "low", "medium"}
	buf := &bytes.Buffer{}
	r := strings.NewReader("y\nn\nn")
	prioritiser.RunCLI(input, r, buf)
}