// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/12.
package solution

import (
	"fmt"
	"strings"
)

type layerPosition struct {
	depth     int
	position  int
	backwards bool
}

func (l layerPosition) Next() layerPosition {
	if l.backwards {
		l.position--
	} else {
		l.position++
	}

	if l.position < 0 {
		l.position = 1
		l.backwards = false
	} else if l.position >= l.depth {
		l.position = l.depth - 2
		l.backwards = true
	}

	return l
}

func Part1(input string) int {
	layers := parse(input)

	maxLayer := 0
	for layer := range layers {
		if layer > maxLayer {
			maxLayer = layer
		}
	}

	layersTracked := map[int]layerPosition{}
	for layer, depth := range layers {
		layersTracked[layer] = layerPosition{depth: depth}
	}

	layersCaughtOn := []int{}
	for i := 0; i <= maxLayer; i++ {
		if position, ok := layersTracked[i]; ok && position.position == 0 {
			layersCaughtOn = append(layersCaughtOn, i)
		}
		for layer, position := range layersTracked {
			layersTracked[layer] = position.Next()
		}
	}

	severity := 0
	for _, layer := range layersCaughtOn {
		severity += layer * layers[layer]
	}

	return severity
}

func Part2(input string) int {
	layers := parse(input)

	maxLayer := 0
	for layer := range layers {
		if layer > maxLayer {
			maxLayer = layer
		}
	}

	answer := make(chan int)
	startingIndex := 0
	startingLayersTracked := map[int]layerPosition{}
	for layer, depth := range layers {
		startingLayersTracked[layer] = layerPosition{depth: depth}
	}

	for {
		if startingIndex != 0 {
			for layer, position := range startingLayersTracked {
				startingLayersTracked[layer] = position.Next()
			}
		}

		layersTracked := map[int]layerPosition{}
		for k, v := range startingLayersTracked {
			layersTracked[k] = v
		}

		go func(startingIndex int) {
			for i := 0; i <= maxLayer; i++ {
				if position, ok := layersTracked[i]; ok && position.position == 0 {
					return
				}
				for layer, position := range layersTracked {
					layersTracked[layer] = position.Next()
				}
			}
			answer <- startingIndex
		}(startingIndex)

		select {
		case a := <-answer:
			return a
		default:
		}

		startingIndex++
	}
}

func parse(input string) map[int]int {
	m := map[int]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var layer, depth int
		fmt.Sscanf(line, "%d: %d", &layer, &depth)
		m[layer] = depth
	}
	return m
}
