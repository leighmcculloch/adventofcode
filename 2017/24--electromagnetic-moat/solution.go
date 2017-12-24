// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/24.
package solution

import (
	"fmt"
	"strings"
)

func Part1(input string) int {
	pm := parsePortMap(input)
	strength := findMaxStrength(pm, 0)
	return strength
}

func findMaxStrength(pm portMap, pins int) (strength int) {
	matchingPorts := pm.Get(pins)
	for _, p := range matchingPorts {
		if p.used {
			continue
		}
		p.used = true

		thisStrength := p.Strength() + findMaxStrength(pm, p.PinsOnOtherSide(pins))
		if thisStrength > strength {
			strength = thisStrength
		}

		p.used = false
	}
	return strength
}

func Part2(input string) int {
	pm := parsePortMap(input)
	_, strength := findMaxLengthMaxStrength(pm, 0)
	return strength
}

func findMaxLengthMaxStrength(pm portMap, pins int) (length, strength int) {
	ports := pm.Get(pins)
	for _, p := range ports {
		if p.used {
			continue
		}
		p.used = true

		childrenLength, childrenStrength := findMaxLengthMaxStrength(pm, p.PinsOnOtherSide(pins))
		thisLength := 1 + childrenLength
		thisStrength := p.Strength() + childrenStrength
		if thisLength > length {
			length = thisLength
			strength = thisStrength
		} else if thisLength == length && thisStrength > strength {
			length = thisLength
			strength = thisStrength
		}

		p.used = false
	}
	return length, strength
}

func parsePortMap(s string) portMap {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	pm := portMap{}
	for _, l := range lines {
		var p port
		fmt.Sscanf(l, "%d/%d", &p.pins[0], &p.pins[1])
		pm.Add(&p)
	}
	return pm
}

type portMap map[int][]*port

func (pm portMap) Add(p *port) {
	for i := 0; i < len(p.pins); i++ {
		ports, exists := pm[p.pins[i]]
		if !exists {
			ports = make([]*port, 0, 1)
		}
		ports = append(ports, p)
		pm[p.pins[i]] = ports
	}
}

func (pm portMap) Get(pins int) []*port {
	return pm[pins]
}

type port struct {
	used bool
	pins [2]int
}

func (p port) Strength() (strength int) {
	for _, pins := range p.pins {
		strength += pins
	}
	return strength
}

func (p port) PinsOnOtherSide(this int) int {
	if p.pins[0] == this {
		return p.pins[1]
	}
	return p.pins[0]
}
