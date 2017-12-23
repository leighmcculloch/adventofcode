// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/23.
package solution

func Part1(input string) (output int) {
	instructions := parse(input)

	p := newProgram(0, instructions, nil, nil)

	defer func() {
		switch e := recover().(type) {
		case nil:
		case programCounterEscapeError:
			output = p.instructionCount["mul"]
		default:
			panic(e)
		}
	}()

	for {
		p.tick()
	}
}
