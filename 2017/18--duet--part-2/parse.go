package solution

import "strings"

func parse(input string) []instruction {
	lines := strings.Split(input, "\n")
	instructions := []instruction{}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			continue
		}
		instr := parseInstruction(lines[i])
		instructions = append(instructions, instr)
	}
	return instructions
}
