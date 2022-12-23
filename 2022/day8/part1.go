package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	matrix, height, width := strToMatrix(input)
	printMatrix(matrix)

	fromLeft := makeMatrix(width, height)
	for h := 0; h < height; h++ {
		fromLeftMax := byte(0)
		for w := 0; w < width; w++ {
			fromLeft[h][w] = fromLeftMax
			th := matrix[h][w]
			if th > fromLeftMax {
				fromLeftMax = th
			}
		}
	}
	printMatrix(fromLeft)

	fromRight := makeMatrix(width, height)
	for h := 0; h < height; h++ {
		fromRightMax := byte(0)
		for w := width - 1; w >= 0; w-- {
			fromRight[h][w] = fromRightMax
			th := matrix[h][w]
			if th > fromRightMax {
				fromRightMax = th
			}
		}
	}
	printMatrix(fromRight)

	fromTop := makeMatrix(width, height)
	for w := 0; w < width; w++ {
		fromTopMax := byte(0)
		for h := 0; h < height; h++ {
			fromTop[h][w] = fromTopMax
			th := matrix[h][w]
			if th > fromTopMax {
				fromTopMax = th
			}
		}
	}
	printMatrix(fromTop)

	fromBottom := makeMatrix(width, height)
	for w := 0; w < width; w++ {
		fromBottomMax := byte(0)
		for h := height - 1; h >= 0; h-- {
			fromBottom[h][w] = fromBottomMax
			th := matrix[h][w]
			if th > fromBottomMax {
				fromBottomMax = th
			}
		}
	}
	printMatrix(fromBottom)

	compares := [][][]byte{fromLeft, fromRight, fromTop, fromBottom}

	count := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			th := matrix[h][w]
			for _, compare := range compares {
				if th > compare[h][w] {
					count++
					break
				}
			}
		}
	}

	fmt.Println("-")
	fmt.Println(count)
}

func makeMatrix(height, width int) [][]byte {
	m := make([][]byte, height)
	for h := 0; h < len(m); h++ {
		m[h] = make([]byte, width)
	}
	return m
}

func strToMatrix(input string) (matrix [][]byte, height, width int) {
	lines := strings.Split(input, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	height = len(lines)
	width = len(lines[0])

	matrix = make([][]byte, height)
	for h, row := range lines {
		matrix[h] = make([]byte, width)
		for w, c := range []byte(row) {
			treeHeight := c - '0' + 1
			matrix[h][w] = treeHeight
		}
	}

	return
}

func printMatrix(matrix [][]byte) {
	fmt.Println("-")
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}
