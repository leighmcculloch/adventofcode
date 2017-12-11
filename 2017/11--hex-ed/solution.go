// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/11.
package solution

import (
	"math"
	"strings"
)

var directions = map[string]func(x, y float64) (float64, float64){
	"n": func(x, y float64) (float64, float64) {
		return x, y + 1
	},
	"ne": func(x, y float64) (float64, float64) {
		return x + 1, y + 0.5
	},
	"se": func(x, y float64) (float64, float64) {
		return x + 1, y - 0.5
	},
	"s": func(x, y float64) (float64, float64) {
		return x, y - 1
	},
	"sw": func(x, y float64) (float64, float64) {
		return x - 1, y - 0.5
	},
	"nw": func(x, y float64) (float64, float64) {
		return x - 1, y + 0.5
	},
}

func Part1(input string) int {
	list := strings.Split(input, ",")
	x := 0.0
	y := 0.0
	for _, direction := range list {
		direction = strings.TrimSpace(direction)
		if direction == "" {
			continue
		}
		directionFunction := directions[direction]
		x, y = directionFunction(x, y)
	}
	if x < 0.0 {
		x = -x
	}
	if y < 0.0 {
		y = -y
	}
	max := x
	if y > max {
		max = y
	}
	max = math.Ceil(max)
	return int(max)
}

func Part2(input string) int {
	list := strings.Split(input, ",")
	x := 0.0
	y := 0.0
	maxX := 0.0
	maxY := 0.0
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
	if maxX < 0 {
		maxX = -maxX
	}
	if maxY < 0 {
		maxY = -maxY
	}
	max := maxX
	if maxY > max {
		max = maxY
	}
	return int(max)
}
