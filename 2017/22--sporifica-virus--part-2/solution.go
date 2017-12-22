// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/22.
package solution

import (
	"strings"
)

func Part2(input string) int {
	g, size := parseGrid(input)
	d := direction(north{})
	p := vector{x: size / 2, y: size / 2}

	causedInfectionsCount := 0

	for i := 0; i < 10000000; i++ {
		switch g.State(p).(type) {
		case clean:
			d = d.Left()
		case weakened:
		case infected:
			d = d.Right()
		case flagged:
			d = d.Reverse()
		}

		g.Evolve(p)
		if _, ok := g.State(p).(infected); ok {
			causedInfectionsCount++
		}

		p = d.Forward(p)
	}

	return causedInfectionsCount
}

type grid map[vector]state

func parseGrid(input string) (g grid, size int) {
	g = grid{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for y, r := range lines {
		for x, c := range r {
			v := vector{x: x, y: y}
			if c == '#' {
				g.Evolve(v)
				g.Evolve(v)
			}
		}
	}
	return g, len(lines)
}

func (g grid) Evolve(v vector) {
	g[v] = g.State(v).Evolve()
}

func (g grid) State(v vector) state {
	if s, ok := g[v]; ok {
		return s
	} else {
		return clean{}
	}
}

type vector struct {
	x, y int
}

type state interface {
	Evolve() state
}

type clean struct{}

func (clean) Evolve() state { return weakened{} }

type weakened struct{}

func (weakened) Evolve() state { return infected{} }

type infected struct{}

func (infected) Evolve() state { return flagged{} }

type flagged struct{}

func (flagged) Evolve() state { return clean{} }

type direction interface {
	Right() direction
	Left() direction
	Reverse() direction
	Forward(vector) vector
}

type north struct{}

func (north) Right() direction        { return east{} }
func (north) Left() direction         { return west{} }
func (north) Reverse() direction      { return south{} }
func (north) Forward(v vector) vector { return vector{x: v.x, y: v.y - 1} }

type east struct{}

func (east) Right() direction        { return south{} }
func (east) Left() direction         { return north{} }
func (east) Reverse() direction      { return west{} }
func (east) Forward(v vector) vector { return vector{x: v.x + 1, y: v.y} }

type south struct{}

func (south) Right() direction        { return west{} }
func (south) Left() direction         { return east{} }
func (south) Reverse() direction      { return north{} }
func (south) Forward(v vector) vector { return vector{x: v.x, y: v.y + 1} }

type west struct{}

func (west) Right() direction        { return north{} }
func (west) Left() direction         { return south{} }
func (west) Reverse() direction      { return east{} }
func (west) Forward(v vector) vector { return vector{x: v.x - 1, y: v.y} }
