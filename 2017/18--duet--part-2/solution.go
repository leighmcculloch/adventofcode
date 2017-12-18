// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/18.
package solution

func Part2(input string) int {
	instructions := parse(input)

	progs := newPrograms()

	to0 := make(chan int, 135)
	to1 := make(chan int, 135)

	progs.add(newProgram(0, instructions, to0, to1))
	progs.add(newProgram(1, instructions, to1, to0))

	for i := 0; !progs.allYielding(); i++ {
		p := progs.get(i % progs.len())
		p.tick()
	}

	return progs.get(1).instructionCount["snd"]
}
