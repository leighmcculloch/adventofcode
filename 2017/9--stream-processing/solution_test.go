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
		{`{}`, 1},
		{`{{{}}}`, 6},
		{`{{},{}}`, 5},
		{`{{{},{},{{}}}}`, 16},
		{`{<a>,<a>,<a>,<a>}`, 1},
		{`{{<ab>},{<ab>},{<ab>},{<ab>}}`, 9},
		{`{{<!!>},{<!!>},{<!!>},{<!!>}}`, 9},
		{`{{<a!>},{<a!>},{<a!>},{<ab>}}`, 3},
		{input, 10820},
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
		{`<>`, 0},
		{`<random characters>`, 17},
		{`<<<<>`, 3},
		{`<{!>}>`, 2},
		{`<!!>`, 0},
		{`<!!!>>`, 0},
		{`<{o"i!a,<{i<a>`, 10},
		{input, 5547},
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
