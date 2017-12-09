// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/9.
package solution

// Part1.
func Part1(input string) int {
	stack := 0
	garbage := false
	ignoreNext := false
	score := 0

	for _, c := range input {
		if ignoreNext {
			ignoreNext = false
			continue
		}
		switch c {
		case '{':
			if !garbage {
				stack++
				score += stack
			}
		case '}':
			if !garbage {
				stack--
			}
		case '<':
			garbage = true
		case '>':
			garbage = false
		case '!':
			ignoreNext = true
		}
	}

	return score
}

// Part2.
func Part2(input string) int {
	garbage := false
	ignoreNext := false
	garbageCount := 0

	for _, c := range input {
		if ignoreNext {
			ignoreNext = false
			continue
		}
		switch c {
		case '<':
			if garbage {
				garbageCount++
			}
			garbage = true
		case '>':
			garbage = false
		case '!':
			ignoreNext = true
		default:
			if garbage {
				garbageCount++
			}
		}
	}

	return garbageCount
}
