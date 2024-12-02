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

// New implementation of isSequenceSafe for part 2
func isSequenceSafe(nums []int) bool {

    // First check if sequence is safe without removing any numbers
    if isSequenceSafeWithoutRemoval(nums) {
        return true
    }

    // Try removing each number one at a time
    for i := range nums {
        withoutCurrent := make([]int, 0, len(nums)-1)
        withoutCurrent = append(withoutCurrent, nums[:i]...)
        withoutCurrent = append(withoutCurrent, nums[i+1:]...)

        if isSequenceSafeWithoutRemoval(withoutCurrent) {
            return true
        }
    }
    return false
}

func isSequenceSafeWithoutRemoval(nums []int) bool {
    if len(nums) < 2 {
        return true
    }

    increasing := nums[1] > nums[0]

    for i := 1; i < len(nums); i++ {
        diff := nums[i] - nums[i-1]

        if diff < -3 || diff > 3 || diff == 0 {
            return false
        }

        if increasing && diff < 0 {
            return false
        }
        if !increasing && diff > 0 {
            return false
        }
    }

    return true
}

