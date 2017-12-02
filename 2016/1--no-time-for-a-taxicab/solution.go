// Package solution contains solutions to the problems described at http://adventofcode.com/2016/day/1.
package solution

import (
	"regexp"
	"strconv"
)

// Part1.
func Part1(input string) int {
	moves := parseMoves(input)

	d := north
	x := 0
	y := 0

	for _, m := range moves {
		d = d.turn(m.turn)
		switch d {
		case north:
			y += m.count
		case south:
			y -= m.count
		case east:
			x += m.count
		case west:
			x -= m.count
		}
	}

	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}

	return x + y
}

// Part2.
func Part2(input string) int {
	return 0
}

type direction string

const (
	north direction = "N"
	south direction = "S"
	east  direction = "E"
	west  direction = "W"
)

func (d direction) turn(t turn) direction {
	switch d {
	case north:
		switch t {
		case right:
			return east
		case left:
			return west
		}
	case south:
		switch t {
		case right:
			return west
		case left:
			return east
		}
	case east:
		switch t {
		case right:
			return south
		case left:
			return north
		}
	case west:
		switch t {
		case right:
			return north
		case left:
			return south
		}
	}
	panic("direction " + string(d) + " or turn " + string(t) + " not recognized")
}

type turn string

const (
	right turn = "R"
	left  turn = "L"
)

type move struct {
	turn  turn
	count int
}

func parseMoves(s string) []move {
	r := regexp.MustCompile("([RL])([0-9]+)")
	matches := r.FindAllStringSubmatch(s, -1)
	moves := make([]move, 0, len(matches))
	for _, m := range matches {
		t := turn(m[1])
		c, err := strconv.Atoi(m[2])
		if err != nil {
			panic(err)
		}
		moves = append(moves, move{t, c})
	}
	return moves
}
