package part1

func Part1(inputs []string) int {
	twos := 0
	threes := 0

	for _, input := range inputs {
		count := map[string]int{}
		for _, b := range input {
			count[string(b)]++
		}
		hasTwo := false
		hasThree := false
		for _, c := range count {
			switch c {
			case 2:
				hasTwo = true
			case 3:
				hasThree = true
			}
		}
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}

	return twos * threes
}
