package solution

import (
	"testing"
)

func TestSolution(t *testing.T) {
	w := 5593
	g := Solution()

	if g == w {
		t.Logf("got %v, want %v", g, w)
	} else {
		t.Errorf("got %v, want %v", g, w)
	}
}
