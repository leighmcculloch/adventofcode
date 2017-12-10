// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/8.
package solution

import (
	"fmt"
	"strings"
)

var operators = map[string]func(lhs, rhs int) bool{
	">": func(lhs, rhs int) bool {
		return lhs > rhs
	},
	"<": func(lhs, rhs int) bool {
		return lhs < rhs
	},
	">=": func(lhs, rhs int) bool {
		return lhs >= rhs
	},
	"<=": func(lhs, rhs int) bool {
		return lhs <= rhs
	},
	"==": func(lhs, rhs int) bool {
		return lhs == rhs
	},
	"!=": func(lhs, rhs int) bool {
		return lhs != rhs
	},
}

var actions = map[string]func(current, change int) int{
	"inc": func(current, change int) int {
		return current + change
	},
	"dec": func(current, change int) int {
		return current - change
	},
}

// Part1.
func Part1(input string) int {
	lines := strings.Split(input, "\n")
	registers := map[string]int{}
	for _, l := range lines {
		if l == "" {
			continue
		}

		var name, action, conditionLHS, conditionOperator string
		var actionValue, conditionRHS int
		fmt.Sscanf(l, "%s %s %d if %s %s %d", &name, &action, &actionValue, &conditionLHS, &conditionOperator, &conditionRHS)

		op := operators[conditionOperator]
		if op(registers[conditionLHS], conditionRHS) {
			registers[name] = actions[action](registers[name], actionValue)
		}
	}

	max := 0
	for _, v := range registers {
		if v > max {
			max = v
		}
	}

	return max
}

// Part2.
func Part2(input string) int {
	lines := strings.Split(input, "\n")
	registers := map[string]int{}
	registerMaxes := map[string]int{}
	for _, l := range lines {
		if l == "" {
			continue
		}

		var name, action, conditionLHS, conditionOperator string
		var actionValue, conditionRHS int
		fmt.Sscanf(l, "%s %s %d if %s %s %d", &name, &action, &actionValue, &conditionLHS, &conditionOperator, &conditionRHS)

		op := operators[conditionOperator]
		if op(registers[conditionLHS], conditionRHS) {
			registers[name] = actions[action](registers[name], actionValue)
			if registers[name] > registerMaxes[name] {
				registerMaxes[name] = registers[name]
			}
		}
	}

	max := 0
	for _, v := range registerMaxes {
		if v > max {
			max = v
		}
	}

	return max
}
