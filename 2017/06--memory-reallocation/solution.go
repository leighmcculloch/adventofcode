// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/6.
package solution

import (
	"fmt"
	"strconv"
	"strings"
)

// Part1.
func Part1(list string) int {
	listNumbers := strings.Fields(list)
	banks := []int{}
	for _, n := range listNumbers {
		if n == "" {
			continue
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		banks = append(banks, i)
	}

	snapshots := map[string]bool{}
	steps := 0
	for {
		steps++

		// find largest
		iOfLargest := 0
		for i := 0; i < len(banks); i++ {
			if banks[i] > banks[iOfLargest] {
				iOfLargest = i
			}
		}

		// circularly iterate around banks redistributing the largest
		valueToRedistribute := banks[iOfLargest]
		banks[iOfLargest] = 0
		for i := 0; i < valueToRedistribute; i++ {
			offsetIndex := (i + iOfLargest + 1) % len(banks)
			banks[offsetIndex]++
		}

		banksStr := fmt.Sprintf("%v", banks)
		if _, ok := snapshots[banksStr]; ok {
			break
		}
		snapshots[banksStr] = true
	}

	return steps
}

// Part2.
func Part2(list string) int {
	listNumbers := strings.Fields(list)
	banks := []int{}
	for _, n := range listNumbers {
		if n == "" {
			continue
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		banks = append(banks, i)
	}

	snapshots := map[string]int{}
	steps := 0
	for {
		steps++

		// find largest
		iOfLargest := 0
		for i := 0; i < len(banks); i++ {
			if banks[i] > banks[iOfLargest] {
				iOfLargest = i
			}
		}

		// circularly iterate around banks redistributing the largest
		valueToRedistribute := banks[iOfLargest]
		banks[iOfLargest] = 0
		for i := 0; i < valueToRedistribute; i++ {
			offsetIndex := (i + iOfLargest + 1) % len(banks)
			banks[offsetIndex]++
		}

		banksStr := fmt.Sprintf("%v", banks)
		if stepFirstSeenAt, ok := snapshots[banksStr]; ok {
			return steps - stepFirstSeenAt
		}
		snapshots[banksStr] = steps
	}
}
