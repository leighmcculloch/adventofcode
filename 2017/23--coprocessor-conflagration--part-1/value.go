package solution

import "strconv"

type value interface {
	Get(registers map[string]int) int
}

func parseValue(s string) value {
	if s == "" {
		return noneValue{}
	}
	if n, err := strconv.Atoi(s); err == nil {
		return fixedValue(n)
	} else {
		return registerValue(s)
	}
}

type registerValue string

func (r registerValue) Get(registers map[string]int) int {
	return registers[string(r)]
}

type fixedValue int

func (f fixedValue) Get(registers map[string]int) int {
	return int(f)
}

type noneValue struct{}

func (n noneValue) Get(registers map[string]int) int {
	return 0
}
