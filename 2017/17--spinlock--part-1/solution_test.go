package solution

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		input      int
		wantOutput int
	}{
		{3, 638},
		{324, 1306},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			s := Part1(c.input)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %v, want %v", g, w)
			} else {
				t.Errorf("got %v, want %v", g, w)
			}
		})
	}
}
