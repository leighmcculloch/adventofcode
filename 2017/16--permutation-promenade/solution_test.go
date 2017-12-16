package solution

import (
	"fmt"
	"io/ioutil"
	"testing"
)

var dance = func() string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}()

func TestPart1(t *testing.T) {
	cases := []struct {
		input      string
		dance      string
		wantOutput string
	}{
		{`abcde`, `s1,x3/4,pe/b`, `baedc`},
		{`abcdefghijklmnop`, dance, `padheomkgjfnblic`},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s-%s...", c.input, c.dance[:6]), func(t *testing.T) {
			s := Part1(c.input, c.dance)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %s, want %s", g, w)
			} else {
				t.Errorf("got %s, want %s", g, w)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		input      string
		dance      string
		wantOutput string
	}{
		{`abcdefghijklmnop`, dance, `bfcdeakhijmlgopn`},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%s-%s...", c.input, c.dance[:6]), func(t *testing.T) {
			s := Part2(c.input, c.dance)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %s, want %s", g, w)
			} else {
				t.Errorf("got %s, want %s", g, w)
			}
		})
	}
}
