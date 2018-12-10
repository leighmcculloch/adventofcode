package part1

func sumOfMetadata(input []int) (consumedCount, metadataSum int) {
	childrenCount := input[0]
	metadataCount := input[1]
	consumedCount += 2
	for i := 0; i < childrenCount; i++ {
		subConsumedCount, subMetadataSum := sumOfMetadata(input[consumedCount:])
		consumedCount += subConsumedCount
		metadataSum += subMetadataSum
	}
	for i := 0; i < metadataCount; i++ {
		metadataSum += input[consumedCount+i]
	}
	consumedCount += metadataCount
	return consumedCount, metadataSum
}
