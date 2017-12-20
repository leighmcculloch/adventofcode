// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/20.
package solution

import (
	"fmt"
	"strings"
)

func Part1(input string) int {
	particles := parse(input)

	for i := 0; i < 1000; i++ {
		for i, p := range particles {
			particles[i] = p.next()
		}
	}

	minIndex := 0
	minDistance := particles[0].distance()
	for i, p := range particles {
		d := p.distance()
		if d < minDistance {
			minDistance = d
			minIndex = i
		}
	}

	return minIndex
}

func parse(input string) []particle {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	particles := make([]particle, 0, len(lines))
	for _, l := range lines {
		var p particle
		fmt.Sscanf(
			l,
			"p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&p.p.x,
			&p.p.y,
			&p.p.z,
			&p.v.x,
			&p.v.y,
			&p.v.z,
			&p.a.x,
			&p.a.y,
			&p.a.z,
		)
		particles = append(particles, p)
	}
	return particles
}

type particle struct {
	p vector
	v vector
	a vector
}

func (p particle) next() particle {
	p.v.x += p.a.x
	p.v.y += p.a.y
	p.v.z += p.a.z
	p.p.x += p.v.x
	p.p.y += p.v.y
	p.p.z += p.v.z
	return p
}

func (p particle) distance() int {
	x := p.p.x
	if x < 0 {
		x = -x
	}
	y := p.p.y
	if y < 0 {
		y = -y
	}
	z := p.p.z
	if z < 0 {
		z = -z
	}
	return x + y + z
}

type vector struct {
	x, y, z int
}
