package part2

func Part2(inputs []string) string {
	for _, input1 := range inputs {
		for _, input2 := range inputs {
			diff := 0
			lastX := 0
			for x := 0; x < len(input1); x++ {
				if input1[x] != input2[x] {
					diff++
					lastX = x
				}
			}
			if diff == 1 {
				letters := ""
				for x, c := range input1 {
					if x != lastX {
						letters += string(c)
					}
				}
				return letters
			}
		}
	}

	return ""
}
