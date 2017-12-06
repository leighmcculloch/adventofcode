package solution

import (
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		input      string
		wantOutput int
	}{
		{`0 2 7 0`, 5},
		{`5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6`, 5042},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := Part1(c.input)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %d, want %d", g, w)
			} else {
				t.Errorf("got %d, want %d", g, w)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		input      string
		wantOutput int
	}{
		{`0 2 7 0`, 4},
		{`5	1	10	0	1	7	13	14	3	12	8	10	7	12	0	6`, 1086},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := Part2(c.input)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %d, want %d", g, w)
			} else {
				t.Errorf("got %d, want %d", g, w)
			}
		})
	}
}
