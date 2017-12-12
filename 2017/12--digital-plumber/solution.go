// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/12.
package solution

import (
	"sort"
	"strings"
)

func Part1(input string) int {
	programs := parse(input)
	connectedToZero := map[string]bool{"0": true}
	previousRunConnectToZeroCount := 0

	firstTime := true
	for firstTime || len(connectedToZero) != previousRunConnectToZeroCount {
		firstTime = false
		previousRunConnectToZeroCount = len(connectedToZero)
		for p, connectedPrograms := range programs {
			for _, c := range connectedPrograms {
				if connectedToZero[c] {
					connectedToZero[p] = true
				}
			}
		}
	}
	return len(connectedToZero)
}

func Part2(input string) int {
	programs := parse(input)

	groups := map[string]bool{}

	for program := range programs {
		connectedToProgram := map[string]bool{program: true}
		previousRunConnectToZeroCount := 0

		firstTime := true
		for firstTime || len(connectedToProgram) != previousRunConnectToZeroCount {
			firstTime = false
			previousRunConnectToZeroCount = len(connectedToProgram)
			for p, connectedPrograms := range programs {
				for _, c := range connectedPrograms {
					if connectedToProgram[c] {
						connectedToProgram[p] = true
					}
				}
			}
		}

		group := []string{}
		for p := range connectedToProgram {
			group = append(group, p)
		}

		sort.Strings(group)
		groupID := strings.Join(group, ",")

		groups[groupID] = true
	}

	return len(groups)
}

func parse(input string) map[string][]string {
	lines := strings.Split(input, "\n")
	programs := map[string][]string{}
	for _, l := range lines {
		if l == "" {
			continue
		}

		parts := strings.Split(l, " <-> ")
		programID := parts[0]
		connectedProgramIDs := strings.Split(parts[1], ", ")

		programs[programID] = connectedProgramIDs
	}
	return programs
}
