// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/15.
package main

import "fmt"

func Part1(startingA, startingB int) int {
	prevA := startingA
	prevB := startingB
	count := 0
	for i := 0; i < 40000000; i++ {
		prevA = prevA * 16807 % 2147483647
		prevB = prevB * 48271 % 2147483647
		if uint16(prevA) == uint16(prevB) {
			count++
		}
	}
	return count
}

func Part2(startingA, startingB int) int {
	matched := 0

	prevA := startingA
	prevB := startingB

	for i := 0; i < 5000000; i++ {
		for {
			prevA = prevA * 16807 % 2147483647
			if prevA%4 == 0 {
				break
			}
		}

		for {
			prevB = prevB * 48271 % 2147483647
			if prevB%8 == 0 {
				break
			}
		}

		if uint16(prevA) == uint16(prevB) {
			matched++
		}
	}

	return matched
}

func main() {
	fmt.Println(Part1(65, 8921))
	fmt.Println(Part2(783, 325))
}
