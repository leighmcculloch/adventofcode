// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/5.
package solution

import (
	"strconv"
	"strings"
)

// Part1.
func Part1(list string) int {
	listLines := strings.Split(list, "\n")
	jumps := []int{}
	for _, l := range listLines {
		if l == "" {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		jumps = append(jumps, n)
	}

	steps := 0
	i := 0
	for i >= 0 && i < len(jumps) {
		v := jumps[i]
		jumps[i]++
		i += v
		steps++
	}

	return steps
}

// Part2.
func Part2(list string) int {
	listLines := strings.Split(list, "\n")
	jumps := []int{}
	for _, l := range listLines {
		if l == "" {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		jumps = append(jumps, n)
	}

	steps := 0
	i := 0
	for i >= 0 && i < len(jumps) {
		v := jumps[i]
		if v >= 3 {
			jumps[i]--
		} else {
			jumps[i]++
		}
		i += v
		steps++
	}

	return steps
}
