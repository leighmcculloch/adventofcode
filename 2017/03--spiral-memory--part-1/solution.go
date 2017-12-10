// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/3.
package solution

import (
	"math"
)

// Part1.
func Part1(index int) int {
	if index == 1 {
		return 0
	}

	sqrt := math.Sqrt(float64(index))
	sqrtUp := int(math.Ceil(sqrt))
	sqrtUpOdd := sqrtUp
	if sqrtUpOdd&1 == 0 {
		sqrtUpOdd++
	}
	worstDistance := sqrtUpOdd - 1
	bestDistance := worstDistance / 2
	distanceRange := worstDistance - bestDistance
	count := sqrtUpOdd*4 - 4
	offsetRatio := 1 - (float64(sqrtUpOdd)-sqrt)/2
	offset := int(offsetRatio * float64(count))

	// http://m.wolframalpha.com/input/?i=plot+abs%28%28x+mod+%282*%284-2%29%29%29+-+%284-2%29%29
	wave := math.Abs(math.Mod(float64(offset), float64(2*distanceRange)) - float64(distanceRange))

	return bestDistance + int(wave)
}
