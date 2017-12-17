package solution

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	cases := []struct {
		input      int
		wantOutput int
	}{
		{324, 20430489},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			s := Part2(c.input)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %v, want %v", g, w)
			} else {
				t.Errorf("got %v, want %v", g, w)
			}
		})
	}
}
