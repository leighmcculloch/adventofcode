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

func TestWalk(t *testing.T) {
	cases := []struct {
		input       string
		wantSteps   int
		wantLetters string
	}{
		{`     |          
     |  +--+    
     A  |  C    
 F---|----E|--+ 
     |  |  |  D 
     +B-+  +--+ `, 38, `ABCDEF`},
		{input, 17450, `XYFDJNRCQA`},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			s, l := Walk(c.input)

			if g, w := s, c.wantSteps; g == w {
				t.Logf("got %v steps, want %v steps", g, w)
			} else {
				t.Errorf("got %v steps, want %v steps", g, w)
			}

			if g, w := l, c.wantLetters; g == w {
				t.Logf("got letters %v, want letters %v", g, w)
			} else {
				t.Errorf("got letters %v, want letters %v", g, w)
			}
		})
	}
}
