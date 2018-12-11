package part2

func getValue(input []int) (consumedCount, value int) {
	childrenCount := input[0]
	metadataCount := input[1]
	consumedCount += 2
	childValues := map[int]int{}
	for i := 0; i < childrenCount; i++ {
		childConsumedCount, childValue := getValue(input[consumedCount:])
		consumedCount += childConsumedCount
		childValues[i] = childValue
	}
	if childrenCount == 0 {
		for i := 0; i < metadataCount; i++ {
			value += input[consumedCount+i]
		}
	} else {
		for i := 0; i < metadataCount; i++ {
			value += childValues[input[consumedCount+i]-1]
		}
	}
	consumedCount += metadataCount
	return consumedCount, value
}
