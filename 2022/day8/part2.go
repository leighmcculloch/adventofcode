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

	score := 0
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			th := matrix[h][w]
			rs := 0
			for frw := w + 1; frw < width; frw++ {
				rs++
				if th <= matrix[h][frw] {
					break
				}
			}
			ls := 0
			for flw := w - 1; flw >= 0; flw-- {
				ls++
				if th <= matrix[h][flw] {
					break
				}
			}
			ds := 0
			for fdh := h + 1; fdh < height; fdh++ {
				ds++
				if th <= matrix[fdh][w] {
					break
				}
			}
			us := 0
			for fuh := h - 1; fuh >= 0; fuh-- {
				us++
				if th <= matrix[fuh][w] {
					break
				}
			}
			m := ls * rs * ds * us
			fmt.Printf("%d,%d %d %d %d %d = %d\n", h, w, us, rs, ds, ls, m)
			if m > score {
				score = m
			}
		}
	}

	fmt.Println("-")
	fmt.Println(score)
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
