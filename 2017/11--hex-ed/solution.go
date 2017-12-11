// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/11.
package solution

import "strings"

var directions = map[string]func(x, y int) (int, int){
	"n": func(x, y int) (int, int) {
		return x, y + 1
	},
	"ne": func(x, y int) (int, int) {
		return x + 1, y + 1
	},
	"se": func(x, y int) (int, int) {
		return x + 1, y - 1
	},
	"s": func(x, y int) (int, int) {
		return x, y - 1
	},
	"sw": func(x, y int) (int, int) {
		return x - 1, y - 1
	},
	"nw": func(x, y int) (int, int) {
		return x - 1, y + 1
	},
}

func Part1(input string) int {
	list := strings.Split(input, ",")
	x := 0
	y := 0
	for _, direction := range list {
		direction = strings.TrimSpace(direction)
		if direction == "" {
			continue
		}
		directionFunction := directions[direction]
		x, y = directionFunction(x, y)
	}
	max := x
	if y > max {
		max = y
	}
	return max
}

func Part2(input string) int {
	list := strings.Split(input, ",")
	x := 0
	y := 0
	maxX := 0
	maxY := 0
	for _, direction := range list {
		direction = strings.TrimSpace(direction)
		if direction == "" {
			continue
		}
		directionFunction := directions[direction]
		x, y = directionFunction(x, y)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}
	max := maxX
	if maxY > max {
		max = maxY
	}
	return max
}
