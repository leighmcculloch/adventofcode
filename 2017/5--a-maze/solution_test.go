package solution

import (
	"io/ioutil"
	"testing"
)

var input = func() string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}()

func TestPart1(t *testing.T) {
	cases := []struct {
		input      string
		wantOutput int
	}{
		{`0
3
0
1
-3`, 5},
		{input, 351282},
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
		{`0
3
0
1
-3`, 10},
		{input, 24568703},
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
