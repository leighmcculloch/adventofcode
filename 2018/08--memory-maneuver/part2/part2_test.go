package part2

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestGetValue(t *testing.T) {
	cases := []struct {
		desc    string
		input   []int
		wantSum int
	}{
		{
			desc:    "part 2: no children",
			input:   inputToIntSlice("0 3 1 1 2"),
			wantSum: 4,
		},
		{
			desc:    "part 2: one child",
			input:   inputToIntSlice("1 3 0 3 10 11 12 1 1 2"),
			wantSum: 66,
		},
		{
			desc:    "part 2: simple example",
			input:   inputToIntSlice("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"),
			wantSum: 66,
		},
		{
			desc: "part 2: input",
			input: inputToIntSlice(func() string {
				b, err := ioutil.ReadFile("../input.txt")
				if err != nil {
					t.Fatalf("error: %v", err)
				}
				return string(b)
			}()),
			wantSum: 33649,
		},
	}

	for _, c := range cases {
		t.Run(c.desc, func(t *testing.T) {
			_, sum := getValue(c.input)
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
