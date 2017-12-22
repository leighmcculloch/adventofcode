// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/22.
package solution

import (
	"strings"
)

func Part1(input string) int {
	g, size := parseGrid(input)
	d := direction(north{})
	p := vector{x: size / 2, y: size / 2}

	causedInfectionsCount := 0

	for i := 0; i < 10000; i++ {
		if g.Infected(p) {
			d = d.Right()
			g.Clean(p)
		} else {
			d = d.Left()
			g.Infect(p)
			causedInfectionsCount++
		}
		p = d.Forward(p)
	}

	return causedInfectionsCount
}

type grid map[vector]bool

func parseGrid(input string) (g grid, size int) {
	g = grid{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for y, r := range lines {
		for x, c := range r {
			v := vector{x: x, y: y}
			if c == '#' {
				g.Infect(v)
			}
		}
	}
	return g, len(lines)
}

func (g grid) Infect(v vector) {
	g[v] = true
}

func (g grid) Clean(v vector) {
	delete(g, v)
}

func (g grid) Infected(v vector) bool {
	return g[v]
}

type vector struct {
	x, y int
}

type direction interface {
	Right() direction
	Left() direction
	Forward(vector) vector
}

type north struct{}

func (north) Right() direction        { return east{} }
func (north) Left() direction         { return west{} }
func (north) Forward(v vector) vector { return vector{x: v.x, y: v.y - 1} }

type east struct{}

func (east) Right() direction        { return south{} }
func (east) Left() direction         { return north{} }
func (east) Forward(v vector) vector { return vector{x: v.x + 1, y: v.y} }

type south struct{}

func (south) Right() direction        { return west{} }
func (south) Left() direction         { return east{} }
func (south) Forward(v vector) vector { return vector{x: v.x, y: v.y + 1} }

type west struct{}

func (west) Right() direction        { return north{} }
func (west) Left() direction         { return south{} }
func (west) Forward(v vector) vector { return vector{x: v.x - 1, y: v.y} }
