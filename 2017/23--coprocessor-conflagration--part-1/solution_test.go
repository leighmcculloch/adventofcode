package solution

import (
	"fmt"
	"testing"
)

func TestPart2(t *testing.T) {
	cases := []struct {
		input      string
		wantOutput int
	}{
		{`
set b 84
set c b
jnz a 2
jnz 1 5
mul b 100
sub b -100000
set c b
sub c -17000
set f 1
set d 2
set e 2
set g d
mul g e
sub g b
jnz g 2
set f 0
sub e -1
set g e
sub g b
jnz g -8
sub d -1
set g d
sub g b
jnz g -13
jnz f 2
sub h -1
set g b
sub g c
jnz g 2
jnz 1 3
sub b -17
jnz 1 -23
`, 6724},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v", c.input), func(t *testing.T) {
			s := Part1(c.input)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %v, want %v", g, w)
			} else {
				t.Errorf("got %v, want %v", g, w)
			}
		})
	}
}
