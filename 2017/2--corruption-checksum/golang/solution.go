// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/2.
package solution

import (
	"strconv"
	"strings"
)

// Part1.
func Part1(spreadsheet string) int {
	matrix := parse(spreadsheet)

	checksum := 0
	for _, r := range matrix {
		min := r[0]
		max := r[0]
		for _, c := range r {
			if c < min {
				min = c
			}
			if c > max {
				max = c
			}
		}
		checksum += max - min
	}
	return checksum
}

// Part2.
func Part2(spreadsheet string) int {
	matrix := parse(spreadsheet)

	checksum := 0
	for _, r := range matrix {
		for ci := 0; ci < len(r); ci++ {
			c := r[ci]
			for ci2 := ci + 1; ci2 < len(r); ci2++ {
				c2 := r[ci2]
				if c%c2 == 0 {
					checksum += c / c2
				} else if c2%c == 0 {
					checksum += c2 / c
				}
			}
		}
	}
	return checksum
}

func parse(spreadsheet string) [][]int {
	rows := strings.Split(spreadsheet, "\n")
	matrix := [][]int{}
	for _, r := range rows {
		cols := []int{}
		for _, c := range strings.Fields(r) {
			n, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			cols = append(cols, n)
		}
		matrix = append(matrix, cols)
	}
	return matrix
}
