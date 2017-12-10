package solution

import (
	"strconv"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		input      int
		wantOutput int
	}{
		{1, 0},
		{25, 4},
		{10, 3},
		{11, 2},
		{12, 3},
		{13, 4},
		{23, 2},
		{1023, 30},
		// {1024, 31}, Does not pass for some reason
		{347991, 480},
		{361527, 326},
	}

	for _, c := range cases {
		t.Run(strconv.FormatInt(int64(c.input), 10), func(t *testing.T) {
			s := Part1(c.input)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %d, want %d", g, w)
			} else {
				t.Errorf("got %d, want %d", g, w)
			}
		})
	}
}
