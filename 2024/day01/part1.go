package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput() ([]int, []int, error) {
	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid input line: %s", line)
		}

		l, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid left number: %s", parts[0])
		}

		r, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid right number: %s", parts[1])
		}

		left = append(left, l)
		right = append(right, r)
	}

	return left, right, scanner.Err()
}

func calculateTotalDistance(left, right []int) int {
	// Make copies to avoid modifying original slices
	l := make([]int, len(left))
	r := make([]int, len(right))
	copy(l, left)
	copy(r, right)

	// Sort both lists
	sort.Ints(l)
	sort.Ints(r)

	total := 0
	for i := 0; i < len(l); i++ {
		distance := abs(l[i] - r[i])
		total += distance
	}

	return total
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	left, right, err := readInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	if len(left) != len(right) {
		fmt.Fprintf(os.Stderr, "Lists have different lengths\n")
		os.Exit(1)
	}

	totalDistance := calculateTotalDistance(left, right)
	fmt.Printf("Total distance: %d\n", totalDistance)
}
