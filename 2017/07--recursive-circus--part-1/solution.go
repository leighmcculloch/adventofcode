// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/7.
package solution

import (
	"fmt"
	"strings"
)

type node struct {
	name   string
	weight int
	parent string
}

func parse(input string) map[string]node {
	lines := strings.Split(input, "\n")

	nodes := map[string]node{}

	for _, l := range lines {
		if l == "" {
			continue
		}

		parts := strings.Split(l, " -> ")

		var name string
		var weight int
		fmt.Sscanf(parts[0], "%s (%d)", &name, &weight)

		var children []string
		if len(parts) > 1 {
			children = strings.Split(parts[1], ", ")
		}

		var parent string
		if n, ok := nodes[name]; ok {
			parent = n.parent
		}
		nodes[name] = node{
			name:   name,
			weight: weight,
			parent: parent,
		}

		for _, c := range children {
			var weight int
			if cn, ok := nodes[c]; ok {
				weight = cn.weight
			}
			nodes[c] = node{
				name:   c,
				weight: weight,
				parent: name,
			}
		}
	}

	return nodes
}

// Part1.
func Part1(input string) string {
	nodes := parse(input)
	for _, n := range nodes {
		if n.parent == "" {
			return n.name
		}
	}
	return ""
}
