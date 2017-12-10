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
		listLength int
		input      string
		wantOutput int
	}{
		{5, `3,4,1,5`, 12},
		{256, input, 29240},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := Part1(c.listLength, c.input)
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
		wantOutput string
	}{
		{``, `a2582a3a0e66e6e86e3812dcb672a272`},
		{`AoC 2017`, `33efeb34ea91902bb2f59c9920caa6cd`},
		{`1,2,3`, `3efbe78a8d82f29979031a4aa0b16a9d`},
		{`1,2,4`, `63960835bcdc130f0b66d7ff4f6a5a8e`},
		{input, `4db3799145278dc9f73dcdbc680bd53d`},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := Part2(c.input)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %s, want %s", g, w)
			} else {
				t.Errorf("got %s, want %s", g, w)
			}
		})
	}
}
