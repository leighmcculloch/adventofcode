// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/10.
package solution

import (
	"fmt"
	"math/bits"
	"strings"
)

// Part1.
func Part1(input string) int {
	used := 0
	for i := 0; i < 128; i++ {
		h := knothash(fmt.Sprintf("%s-%d", input, i))
		for _, b := range h {
			used += bits.OnesCount8(b)
		}
	}
	return used
}

// Part2.
func Part2(input string) int {
	grid := [128][128]bool{}

	for r := range grid {
		h := knothash(fmt.Sprintf("%s-%d", input, r))
		for c1, b := range h {
			grid[r][c1*8+0] = b&(1<<7) != 0
			grid[r][c1*8+1] = b&(1<<6) != 0
			grid[r][c1*8+2] = b&(1<<5) != 0
			grid[r][c1*8+3] = b&(1<<4) != 0
			grid[r][c1*8+4] = b&(1<<3) != 0
			grid[r][c1*8+5] = b&(1<<2) != 0
			grid[r][c1*8+6] = b&(1<<1) != 0
			grid[r][c1*8+7] = b&(1<<0) != 0
		}
	}

	regions := regions{}
	for r := 0; r < 128; r++ {
		for c := 0; c < 128; c++ {
			if !grid[r][c] {
				continue
			}
			rc := coord{r, c}
			if regions.includes(rc) {
				continue
			}

			region := newRegion()
			region.add(rc)

			for i := 0; i < region.len(); i++ {
				potentials := region.get(i).neighbors()
				for _, p := range potentials {
					if grid[p.x][p.y] && !region.includes(p) {
						region.add(p)
					}
				}
			}

			regions = append(regions, region)
		}
	}

	return len(regions)
}

type regions []region

func (rs regions) includes(c coord) bool {
	for _, r := range rs {
		if r.includes(c) {
			return true
		}
	}
	return false
}

type region struct {
	coords       map[coord]bool
	coordsSorted []coord
}

func newRegion() region {
	return region{
		coords:       map[coord]bool{},
		coordsSorted: []coord{},
	}
}

func (r *region) add(c coord) {
	r.coords[c] = true
	r.coordsSorted = append(r.coordsSorted, c)
}

func (r region) includes(c coord) bool {
	return r.coords[c]
}

func (r region) len() int {
	return len(r.coordsSorted)
}

func (r region) get(i int) coord {
	return r.coordsSorted[i]
}

type coord struct {
	x, y int
}

func (c coord) neighbors() []coord {
	coords := []coord{}
	if c.x != 0 {
		coords = append(coords, coord{c.x - 1, c.y})
	}
	if c.x != 127 {
		coords = append(coords, coord{c.x + 1, c.y})
	}
	if c.y != 0 {
		coords = append(coords, coord{c.x, c.y - 1})
	}
	if c.y != 127 {
		coords = append(coords, coord{c.x, c.y + 1})
	}
	return coords
}

func knothash(input string) [16]byte {
	input = strings.TrimSpace(input)
	lengths := append([]byte(input), 17, 31, 73, 47, 23)

	list := [256]byte{}
	for i := range list {
		list[i] = byte(i)
	}

	pos := 0
	skip := 0

	for c := 0; c < 64; c++ {
		for _, l := range lengths {
			for i := pos; i < pos+(int(l)/2); i++ {
				ci := i % len(list)
				si := (pos + int(l) - 1 - (i - pos)) % len(list)
				list[ci], list[si] = list[si], list[ci]
			}
			pos += int(l) + skip
			skip++
		}
	}

	result := [16]byte{}
	for i := 0; i < len(result); i++ {
		for j := 0; j < 16; j++ {
			result[i] ^= list[i*16+j]
		}
	}

	return result
}
