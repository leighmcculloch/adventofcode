// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/20.
package solution

import (
	"fmt"
	"strings"
)

func Part2(input string) int {
	particles := parse(input)

	for t := 0; t < 2000; t++ {
		positions := map[vector]int{}
		for i, p := range particles {
			if p.collided {
				continue
			}
			particles[i] = p.next()
			if o, exists := positions[p.p]; exists {
				particles[i].collided = true
				particles[o].collided = true
			} else {
				positions[p.p] = i
			}
		}

	}

	particlesAlive := 0
	for _, p := range particles {
		if !p.collided {
			particlesAlive++
		}
	}
	return particlesAlive
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
	collided bool
	p        vector
	v        vector
	a        vector
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
