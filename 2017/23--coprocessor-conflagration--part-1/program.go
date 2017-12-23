package solution

import "fmt"

type programCounterEscapeError struct {
	program *program
}

func (e programCounterEscapeError) Error() string {
	return fmt.Sprintf(
		"Program %d has escaped it's instructions, with program counter %d but max instruction %d.",
		e.program.id,
		e.program.programCounter,
		len(e.program.instructions),
	)
}

type program struct {
	id               int
	instructions     []instruction
	instructionCount map[string]int
	programCounter   int
	registers        map[string]int
	in               <-chan int
	out              chan<- int
	yielding         bool
}

func newProgram(id int, instructions []instruction, in <-chan int, out chan<- int) *program {
	return &program{
		id: id,
		registers: map[string]int{
			"p": id,
		},
		instructions:     instructions,
		instructionCount: map[string]int{},
		in:               in,
		out:              out,
	}
}

func (p *program) tick() {
	if p.programCounter >= len(p.instructions) {
		panic(programCounterEscapeError{program: p})
	}

	i := p.instructions[p.programCounter]

	jump, yield := p.execute(i)

	p.programCounter += jump
	p.yielding = yield
}

func (p *program) execute(i instruction) (jump int, yield bool) {
	xv := i.x.Get(p.registers)
	yv := i.y.Get(p.registers)

	var skipped bool
	xv, jump, skipped, yield = instructionSet[i.name](xv, yv, p.in, p.out)

	switch x := i.x.(type) {
	case registerValue:
		p.registers[string(x)] = xv
	}

	if !skipped {
		p.instructionCount[i.name]++
	}

	return jump, yield
}
