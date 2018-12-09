package main

import (
	"fmt"
	"testing"
)

func TestGetWinningScore(t *testing.T) {
	cases := []struct {
		description      string
		playerCount      int
		lastMarble       int
		wantWinningScore int
	}{
		{
			description:      "part1: simple example",
			playerCount:      9,
			lastMarble:       25,
			wantWinningScore: 32,
		},
		{
			description:      "part1: example 1",
			playerCount:      10,
			lastMarble:       1618,
			wantWinningScore: 8317,
		},
		{
			description:      "part1: example 2",
			playerCount:      13,
			lastMarble:       7999,
			wantWinningScore: 146373,
		},
		{
			description:      "part1: example 3",
			playerCount:      17,
			lastMarble:       1104,
			wantWinningScore: 2764,
		},
		{
			description:      "part1: example 4",
			playerCount:      21,
			lastMarble:       6111,
			wantWinningScore: 54718,
		},
		{
			description:      "part1: example 5",
			playerCount:      30,
			lastMarble:       5807,
			wantWinningScore: 37305,
		},
		{
			description:      "part1",
			playerCount:      418,
			lastMarble:       71339,
			wantWinningScore: 412127,
		},
		{
			description:      "part2",
			playerCount:      418,
			lastMarble:       7133900,
			wantWinningScore: 3482394794,
		},
	}

	for _, c := range cases {
		d := fmt.Sprintf("%s: %d, %d = %d", c.description, c.playerCount, c.lastMarble, c.wantWinningScore)
		t.Run(d, func(t *testing.T) {
			winningScore := getWinningScore(c.playerCount, c.lastMarble)
			if winningScore == c.wantWinningScore {
				t.Logf("got %d", winningScore)
			} else {
				t.Errorf("got %d, want %d", winningScore, c.wantWinningScore)
			}
		})
	}
}
