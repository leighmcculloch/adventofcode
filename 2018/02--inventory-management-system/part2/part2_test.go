package part2

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestPart2Example(t *testing.T) {
	inputs := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	letters := Part2(inputs)

	if letters != "fgij" {
		t.Errorf("got letters %q; want %q", letters, 1)
	}
}

func TestPart2(t *testing.T) {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	inputs := strings.Split(strings.TrimSpace(string(b)), "\n")

	letters := Part2(inputs)

	if letters != "cypueihajytordkgzxfqplbwn" {
		t.Errorf("got letters %q; want %q", letters, 1)
	}
}
