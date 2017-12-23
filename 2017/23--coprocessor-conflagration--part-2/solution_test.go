package solution

import (
	"testing"
)

func TestPart2(t *testing.T) {
	w := 903
	g := Part2()
	if g == w {
		t.Logf("got %v, want %v", g, w)
	} else {
		t.Errorf("got %v, want %v", g, w)
	}
}
