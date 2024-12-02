package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	safeCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := parseNumbers(scanner.Text())
		if isSequenceSafe(numbers) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

func parseNumbers(line string) []int {
	fields := strings.Fields(line)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		num, _ := strconv.Atoi(field)
		numbers[i] = num
	}
	return numbers
}

func isSequenceSafe(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	// Determine if sequence is increasing or decreasing
	increasing := nums[1] > nums[0]

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		// Check if difference is between 1 and 3
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		// Check if sequence maintains direction
		if increasing && diff < 0 {
			return false
		}
		if !increasing && diff > 0 {
			return false
		}
	}

	return true
}
