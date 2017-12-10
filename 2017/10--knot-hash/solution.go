// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/9.
package solution

import (
	"encoding/hex"
	"strconv"
	"strings"
)

// Part1.
func Part1(listLength int, input string) int {
	lengths := []int{}
	for _, l := range strings.Split(input, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(l))
		if err != nil {
			panic(err)
		}
		lengths = append(lengths, n)
	}

	list := make([]int, listLength)
	for i := range list {
		list[i] = i
	}

	pos := 0
	skip := 0

	for _, l := range lengths {
		if l > len(list) {
			continue
		}
		for i := pos; i < pos+(l/2); i++ {
			ci := i % len(list)
			si := (pos + l - 1 - (i - pos)) % len(list)
			list[ci], list[si] = list[si], list[ci]
		}
		pos += l + skip
		skip++
	}

	return list[0] * list[1]
}

// Part2.
func Part2(input string) string {
	input = strings.TrimSpace(input)
	lengths := append([]byte(input), 17, 31, 73, 47, 23)

	list := [256]byte{}
	for i := range list {
		list[i] = byte(i)
	}

	pos := 0
	skip := 0

	for c := 0; c < 64; c++ {
		for _, l := range lengths {
			for i := pos; i < pos+(int(l)/2); i++ {
				ci := i % len(list)
				si := (pos + int(l) - 1 - (i - pos)) % len(list)
				list[ci], list[si] = list[si], list[ci]
			}
			pos += int(l) + skip
			skip++
		}
	}

	result := [16]byte{}
	for i := 0; i < len(result); i++ {
		for j := 0; j < 16; j++ {
			result[i] ^= list[i*16+j]
		}
	}

	return hex.EncodeToString(result[:])
}
