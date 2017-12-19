// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/19.
package solution

import (
	"strings"
	"unicode"
)

func Walk(input string) (steps int, letters string) {
	g := parseGrid(input)

	pos := xy{}

	for {
		if g.get(pos).set {
			break
		}
		pos.x++
	}
	steps++

	d := direction(south{})

	for {
		nd := direction(d)
		npos := d.Next(pos)
		s := g.get(npos)
		if !s.set {
			nd = d.Right()
			npos = nd.Next(pos)
			s = g.get(npos)
			if !s.set {
				nd = d.Left()
				npos = nd.Next(pos)
				s = g.get(npos)
				if !s.set {
					return steps, letters
				}
			}
		}

		if s.hasLetter() {
			letters += string(s.letter)
		}

		pos, d = npos, nd
		steps++
	}
}

type grid [][]square

func parseGrid(input string) grid {
	lines := strings.Split(input, "\n")
	g := newGrid(len(lines[0]), len(lines))
	for y, r := range lines {
		for x, c := range r {
			if c != ' ' {
				if unicode.IsLetter(c) {
					g.set(xy{x: x, y: y}, square{set: true, letter: c})
				} else {
					g.set(xy{x: x, y: y}, square{set: true})
				}
			}
		}
	}
	return g
}

func newGrid(width, height int) grid {
	g := make([][]square, height)
	for i := range g {
		g[i] = make([]square, width)
	}
	return g
}

func (g grid) width() int {
	return len(g[0])
}

func (g grid) height() int {
	return len(g)
}

func (g grid) set(pos xy, s square) {
	g[pos.y][pos.x] = s
}

func (g grid) get(pos xy) square {
	if pos.x < 0 || pos.x >= g.width() {
		return square{}
	}
	if pos.y < 0 || pos.y >= g.height() {
		return square{}
	}
	return g[pos.y][pos.x]
}

func (g grid) String() string {
	s := ""
	for y := 0; y < g.height(); y++ {
		for x := 0; x < g.width(); x++ {
			b := g.get(xy{x: x, y: y})
			switch {
			case b.hasLetter():
				s += string([]rune{b.letter})
			case b.set:
				s += "."
			default:
				s += " "
			}
		}
		s += "\n"
	}
	return s
}

type xy struct {
	x, y int
}

type square struct {
	set    bool
	letter rune
}

func (s square) hasLetter() bool {
	var zeroLetter rune
	return s.letter != zeroLetter
}

type direction interface {
	Right() direction
	Left() direction
	Next(xy) xy
}

type north struct{}

func (north) Right() direction { return east{} }
func (north) Left() direction  { return west{} }
func (north) Next(pos xy) xy   { return xy{x: pos.x, y: pos.y - 1} }

type east struct{}

func (east) Right() direction { return south{} }
func (east) Left() direction  { return north{} }
func (east) Next(pos xy) xy   { return xy{x: pos.x + 1, y: pos.y} }

type south struct{}

func (south) Right() direction { return west{} }
func (south) Left() direction  { return east{} }
func (south) Next(pos xy) xy   { return xy{x: pos.x, y: pos.y + 1} }

type west struct{}

func (west) Right() direction { return north{} }
func (west) Left() direction  { return south{} }
func (west) Next(pos xy) xy   { return xy{x: pos.x - 1, y: pos.y} }
