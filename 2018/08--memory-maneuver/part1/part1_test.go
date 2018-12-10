package part1

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestSumOfMetadata(t *testing.T) {
	cases := []struct {
		desc    string
		input   []int
		wantSum int
	}{
		{
			desc:    "part 1: no children",
			input:   inputToIntSlice("0 3 1 1 2"),
			wantSum: 4,
		},
		{
			desc:    "part 1: one child",
			input:   inputToIntSlice("1 3 0 3 10 11 12 1 1 2"),
			wantSum: 37,
		},
		{
			desc:    "part 1: simple example",
			input:   inputToIntSlice("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"),
			wantSum: 138,
		},
		{
			desc: "part 1: input",
			input: inputToIntSlice(func() string {
				b, err := ioutil.ReadFile("../input.txt")
				if err != nil {
					t.Fatalf("error: %v", err)
				}
				return string(b)
			}()),
			wantSum: 42196,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			_, sum := sumOfMetadata(c.input)
			if sum == c.wantSum {
				t.Logf("got %d", sum)
			} else {
				t.Errorf("got %d, want %d", sum, c.wantSum)
			}
		})
	}
}

func inputToIntSlice(input string) []int {
	parts := strings.Split(input, " ")
	intParts := make([]int, len(parts))
	for i, p := range parts {
		pv, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			panic(err)
		}
		intParts[i] = pv
	}
	return intParts
}
