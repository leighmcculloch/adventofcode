package solution

import (
	"strconv"
	"testing"
)

func TestPart2(t *testing.T) {
	cases := []struct {
		input      int
		wantOutput int
	}{
		{746, 747},
		{747, 806},
		{748, 806},
		{749, 806},
		{805, 806},
		{347991, 349975},
	}

	for _, c := range cases {
		t.Run(strconv.FormatInt(int64(c.input), 10), func(t *testing.T) {
			s := Part2(c.input)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %d, want %d", g, w)
			} else {
				t.Errorf("got %d, want %d", g, w)
			}
		})
	}
}
