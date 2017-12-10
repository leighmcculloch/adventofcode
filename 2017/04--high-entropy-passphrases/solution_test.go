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
		{`aa bb cc dd ee`, 1},
		{`aa bb cc dd aa`, 0},
		{`aa bb cc dd aaa`, 1},
		{`aa bb cc dd ee
		  aa bb cc dd aa
		  aa bb cc dd aaa`, 2},
		{input, 383},
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
		{`abcde fghij`, 1},
		{`abcde xyz ecdab`, 0},
		{`a ab abc abd abf abj`, 1},
		{`iiii oiii ooii oooi oooo`, 1},
		{`oiii ioii iioi iiio`, 0},
		{`abcde fghij
		  abcde xyz ecdab
		  a ab abc abd abf abj
		  iiii oiii ooii oooi oooo
		  oiii ioii iioi iiio`, 3},
		{input, 265},
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
