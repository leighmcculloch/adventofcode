// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/1.
package solution

// Part1.
func Part1(digits string) int {
	count := len(digits)
	sum := 0
	for i := 0; i < count; i++ {
		d := digits[i]

		ni := i + 1
		if ni == count {
			ni = 0
		}
		nd := digits[ni]

		if d == nd {
			sum += int(int(d) - '0')
		}
	}
	return sum
}

// Part2.
func Part2(digits string) int {
	count := len(digits)
	sum := 0
	for i := 0; i < count; i++ {
		d := digits[i]

		ni := (i + count/2) % count
		nd := digits[ni]

		if d == nd {
			sum += int(int(d) - '0')
		}
	}
	return sum
}
