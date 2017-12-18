package solution

import "strings"

type instruction struct {
	name string
	x    value
	y    value
}

func parseInstruction(s string) instruction {
	args := [3]string{}
	copy(args[:], strings.Fields(s))
	return instruction{
		name: args[0],
		x:    parseValue(args[1]),
		y:    parseValue(args[2]),
	}
}
