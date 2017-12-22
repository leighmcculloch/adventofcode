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
		{`..#
#..
...`, 5587},
		{input, 5261},
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
