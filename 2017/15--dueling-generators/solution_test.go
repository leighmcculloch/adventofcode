package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	cases := []struct {
		startingA  int
		startingB  int
		wantOutput int
	}{
		{65, 8921, 588},
		{783, 325, 650},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d,%d", c.startingA, c.startingB), func(t *testing.T) {
			s := Part1(c.startingA, c.startingB)
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
		startingA  int
		startingB  int
		wantOutput int
	}{
		{65, 8921, 309},
		{783, 325, 336},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d,%d", c.startingA, c.startingB), func(t *testing.T) {
			s := Part2(c.startingA, c.startingB)
			if g, w := s, c.wantOutput; g == w {
				t.Logf("got %d, want %d", g, w)
			} else {
				t.Errorf("got %d, want %d", g, w)
			}
		})
	}
}
