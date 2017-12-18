package solution

type programs []*program

func newPrograms() programs {
	return programs{}
}

func (ps programs) len() int {
	return len(ps)
}

func (ps programs) get(i int) *program {
	return ps[i]
}

func (ps *programs) add(p *program) {
	*ps = append(*ps, p)
}

func (ps programs) allYielding() bool {
	for _, p := range ps {
		if !p.yielding {
			return false
		}
	}
	return true
}
