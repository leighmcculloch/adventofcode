package solution

import (
	"fmt"
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
		{`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`, 31},
		{input, 1859},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			o := Part1(c.input)

			if g, w := o, c.wantOutput; g == w {
				t.Logf("got %v, want %v", g, w)
			} else {
				t.Errorf("got %v, want %v", g, w)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		input      string
		wantOutput int
	}{
		{`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`, 19},
		{input, 1799},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			o := Part2(c.input)

			if g, w := o, c.wantOutput; g == w {
				t.Logf("got %v, want %v", g, w)
			} else {
				t.Errorf("got %v, want %v", g, w)
			}
		})
	}
}
