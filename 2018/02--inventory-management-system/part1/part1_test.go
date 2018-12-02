package part1

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	inputs := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	checksum := Part1(inputs)

	if checksum != 12 {
		t.Errorf("got checksum %d; want %d", checksum, 1)
	}
}

func TestPart1(t *testing.T) {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	inputs := strings.Split(string(b), "\n")

	checksum := Part1(inputs)

	if checksum != 5166 {
		t.Errorf("got checksum %d; want %d", checksum, 1)
	}
}
