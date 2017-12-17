// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/17.
package solution

func Part1(steps int) int {
	c := newCircbuf()
	c.Insert(0)
	for i := 1; i <= 2017; i++ {
		c.Step(steps)
		c.Insert(i)
	}
	c.Step(1)
	return c.Get()
}

type circbuf struct {
	buf []int
	i   int
}

func newCircbuf() circbuf {
	return circbuf{buf: []int{}}
}

func (c *circbuf) Step(steps int) {
	if len(c.buf) == 0 {
		return
	}
	c.i = (c.i + steps) % len(c.buf)
}

func (c *circbuf) Insert(v int) {
	c.buf = append(c.buf, 0)
	c.Step(1)
	if c.i < len(c.buf) {
		copy(c.buf[c.i+1:], c.buf[c.i:])
	}
	c.buf[c.i] = v
}

func (c circbuf) Get() int {
	return c.buf[c.i]
}
