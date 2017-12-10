// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/3.
package solution

// Part2.
func Part2(min int) int {
	s1 := square{
		Coord: coord{0, 0},
		Value: 1,
	}

	squares := map[coord]square{
		s1.Coord: s1,
	}

	d := east
	last := s1
	for last.Value <= min {
		md := d.Next()
		mc := last.Coord.Add(md)
		if _, ok := squares[mc]; !ok {
			d = md
		}

		s := square{
			Coord: last.Coord.Add(d),
			Value: 0,
		}
		for _, a := range s.Coord.Adjacent() {
			s.Value += squares[a].Value
		}

		squares[s.Coord] = s

		last = s
	}

	return last.Value
}

const (
	east  direction = "E"
	north direction = "N"
	west  direction = "W"
	south direction = "S"
)

type direction string

func (d direction) Next() direction {
	switch d {
	case east:
		return north
	case north:
		return west
	case west:
		return south
	case south:
		return east
	}
	panic("unrecognized direction")
}

func (d direction) Velocity() (x int, y int) {
	switch d {
	case east:
		return 1, 0
	case north:
		return 0, 1
	case west:
		return -1, 0
	case south:
		return 0, -1
	}
	panic("unrecognized direction")
}

type coord struct {
	X, Y int
}

func (c coord) Adjacent() [8]coord {
	return [8]coord{
		{c.X + 1, c.Y},
		{c.X + 1, c.Y + 1},
		{c.X + 1, c.Y - 1},
		{c.X - 1, c.Y},
		{c.X - 1, c.Y + 1},
		{c.X - 1, c.Y - 1},
		{c.X, c.Y + 1},
		{c.X, c.Y - 1},
	}
}

func (c coord) Add(d direction) coord {
	vx, vy := d.Velocity()
	return coord{c.X + vx, c.Y + vy}
}

type square struct {
	Coord coord
	Value int
}
