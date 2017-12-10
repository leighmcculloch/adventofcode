// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/7.
package solution

import (
	"fmt"
	"strings"
)

type node struct {
	name        string
	weight      int
	totalWeight int
	parent      *node
	children    map[string]*node
}

func parse(input string) map[string]*node {
	lines := strings.Split(input, "\n")

	nodes := map[string]*node{}

	for _, l := range lines {
		if l == "" {
			continue
		}

		parts := strings.Split(l, " -> ")

		var name string
		var weight int
		fmt.Sscanf(parts[0], "%s (%d)", &name, &weight)

		var childrenNames []string
		if len(parts) > 1 {
			childrenNames = strings.Split(parts[1], ", ")
		}

		if _, ok := nodes[name]; !ok {
			nodes[name] = &node{name: name}
		}
		nodes[name].weight = weight

		for _, cn := range childrenNames {
			if _, ok := nodes[cn]; !ok {
				nodes[cn] = &node{name: cn}
			}
			nodes[cn].parent = nodes[name]
			if nodes[name].children == nil {
				nodes[name].children = map[string]*node{}
			}
			nodes[name].children[cn] = nodes[cn]
		}
	}

	return nodes
}

func fillTotalWeight(n *node) {
	n.totalWeight = n.weight
	for _, c := range n.children {
		fillTotalWeight(c)
		n.totalWeight += c.totalWeight
	}
}

func find(n *node) *int {
	weights := map[int][]*node{}
	for _, c := range n.children {
		v := find(c)
		if v != nil {
			return v
		}
		if weights[c.totalWeight] == nil {
			weights[c.totalWeight] = []*node{}
		}
		weights[c.totalWeight] = append(weights[c.totalWeight], c)
	}

	if len(weights) <= 1 {
		return nil
	}

	w1 := 0
	n1 := (*node)(nil)
	for w, nodes := range weights {
		if len(nodes) == 1 {
			w1 = w
			n1 = nodes[0]
			break
		}
	}
	w2 := 0
	for w := range weights {
		if w != w1 {
			w2 = w
			break
		}
	}

	diff := n1.weight + (w2 - w1)
	return &diff
}

// Part1.
func Part2(input string) int {
	nodes := parse(input)

	var rootName string
	for _, n := range nodes {
		if n.parent == nil {
			rootName = n.name
		}
	}

	fillTotalWeight(nodes[rootName])

	value := find(nodes[rootName])

	return *value
}
