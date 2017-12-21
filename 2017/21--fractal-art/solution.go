// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/21.
package solution

import (
	"math"
	"strings"
)

func Part1(input string) int {
	return iterate(input, 5)
}

func Part2(input string) int {
	return iterate(input, 18)
}

func iterate(input string, iterations int) int {
	rules := parse(input)

	g := grid(".#./..#/###")

	for i := 0; i < iterations; i++ {
		grids := g.Split()
		for g := range grids {
			for _, v := range grids[g].AllVariations() {
				if m, ok := rules[v]; ok {
					grids[g] = m
					break
				}
			}
		}
		g = join(grids)
	}

	return g.On()
}

func parse(input string) map[grid]grid {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rules := map[grid]grid{}
	for _, l := range lines {
		if l == "" {
			continue
		}
		parts := strings.Split(l, " => ")
		rules[grid(parts[0])] = grid(parts[1])
	}
	return rules
}

type grid string

func join(grids []grid) grid {
	gridSize := grids[0].Size()
	gridsGridWidth := int(math.Sqrt(float64(len(grids))))
	size := gridsGridWidth * gridSize
	joined := make([]byte, 0, size*size+size)
	for h := 0; h < size; h++ {
		for w := 0; w < size; w++ {
			gridX := w / gridSize
			gridY := h / gridSize
			gridIndex := gridX + gridY*gridsGridWidth
			g := grids[gridIndex]
			v := g.Get(w%g.Size(), h%g.Size())
			joined = append(joined, v)
			if (w+1)%size == 0 {
				joined = append(joined, '/')
			}
		}
	}
	return grid(joined)
}

func (g grid) Size() int {
	return strings.Index(string(g), "/")
}

func (g grid) Get(x, y int) byte {
	p := y*g.Size() + y + x
	return g[p]
}

func (g grid) IsOn(x, y int) bool {
	return g.Get(x, y) == '#'
}

func (g grid) On() int {
	on := 0
	for i := 0; i < len(g); i++ {
		if g[i] == '#' {
			on++
		}
	}
	return on
}

func (g grid) Split() []grid {
	var size int

	if g.Size()%2 == 0 {
		size = 2
	} else {
		size = 3
	}

	gridWidth := g.Size() / size
	numGrids := gridWidth * gridWidth
	grids := make([]grid, numGrids)
	for h := 0; h < g.Size(); h++ {
		for w := 0; w < g.Size(); w++ {
			gridX := w / size
			gridY := h / size
			gridIndex := gridX + gridY*gridWidth
			grids[gridIndex] += grid(g.Get(w, h))
			if (w+1)%size == 0 {
				grids[gridIndex] += grid("/")
			}
		}
	}

	return grids
}

func (g grid) AllVariations() []grid {
	r1 := g.RotateRight()
	r2 := r1.RotateRight()
	r3 := r2.RotateRight()
	return []grid{
		g.FlipHorizontal(),
		g.FlipVertical(),

		r1,
		r1.FlipHorizontal(),
		r1.FlipVertical(),

		r2,
		r2.FlipHorizontal(),
		r2.FlipVertical(),

		r3,
		r3.FlipHorizontal(),
		r3.FlipVertical(),
	}
}

func (g grid) FlipHorizontal() grid {
	rows := strings.Split(string(g), "/")
	for r := 0; r < len(rows); r++ {
		b := []byte(rows[r])
		for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
			b[l], b[r] = b[r], b[l]
		}
		rows[r] = string(b)
	}
	return grid(strings.Join(rows, "/"))
}

func (g grid) FlipVertical() grid {
	rows := strings.Split(string(g), "/")
	for l, r := 0, len(rows)-1; l < r; l, r = l+1, r-1 {
		rows[l], rows[r] = rows[r], rows[l]
	}
	return grid(strings.Join(rows, "/"))
}

func (g grid) RotateRight() grid {
	r := ""
	for w := 0; w < g.Size(); w++ {
		for h := 0; h < g.Size(); h++ {
			r += string(g.Get(w, g.Size()-h-1))
		}
		r += "/"
	}
	return grid(r[:len(r)-1])
}

func (g grid) String() string {
	s := ""
	for h := 0; h < g.Size(); h++ {
		for w := 0; w < g.Size(); w++ {
			s += string(g.Get(w, h))
		}
		s += "\n"
	}
	return s
}
