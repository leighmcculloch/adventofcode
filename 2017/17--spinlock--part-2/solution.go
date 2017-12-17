// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/17.
package solution

func Part2(steps int) int {
	c := circbuf{}
	c.Insert(0)
	for i := 1; i <= 50000000; i++ {
		c.Step(steps)
		c.Insert(i)
	}
	return c.GetValue1()
}

type circbuf struct {
	len    int
	i      int
	value1 int
}

func (c *circbuf) Step(steps int) {
	if c.len == 0 {
		return
	}
	c.i = (c.i + steps) % c.len
}

func (c *circbuf) Insert(v int) {
	c.len++
	c.Step(1)
	if c.i == 1 {
		c.value1 = v
	}
}

func (c circbuf) GetValue1() int {
	return c.value1
}
