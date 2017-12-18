// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/18.
package solution

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")

	regs := map[string]int{}
	instructions := map[string]func(x int, y int) (newX int, increment int, terminate bool){
		"set": set,
		"add": add,
		"mul": mul,
		"mod": mod,
		"snd": snd,
		"rcv": rcv,
		"jgz": jgz,
	}

	for i := 0; i < len(lines); {
		l := lines[i]
		if l == "" {
			i++
			continue
		}
		parts := strings.Fields(l)
		instr := instructions[parts[0]]
		reg := parts[1]
		var x = regs[reg]
		var y int
		if len(parts) > 2 {
			var err error
			y, err = strconv.Atoi(parts[2])
			if err != nil {
				y = regs[parts[2]]
			}
		}
		newX, increment, terminate := instr(x, y)
		regs[reg] = newX
		if terminate {
			return newX
		}
		i += increment
	}
	return 0
}

func set(x int, y int) (int, int, bool) {
	return y, 1, false
}

func add(x int, y int) (int, int, bool) {
	return x + y, 1, false
}

func mul(x int, y int) (int, int, bool) {
	return x * y, 1, false
}

func mod(x int, y int) (int, int, bool) {
	return x % y, 1, false
}

var freq = 0

func snd(x int, y int) (int, int, bool) {
	freq = x
	return x, 1, false
}

func rcv(x int, y int) (int, int, bool) {
	if x != 0 {
		return freq, 1, true
	}
	return x, 1, false
}

func jgz(x int, y int) (int, int, bool) {
	if x > 0 {
		return x, y, false
	} else {
		return x, 1, false
	}
}
