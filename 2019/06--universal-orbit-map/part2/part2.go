package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var inputs = func() []string {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	inputs := []string{}
	for _, l := range strings.Split(string(b), "\n") {
		if l == "" {
			continue
		}
		inputs = append(inputs, l)
	}
	return inputs
}()

type Node struct {
	Name   string
	Orbits []*Node
}

type Nodes struct {
	NodeMap map[string]*Node
}

func (ns *Nodes) init() {
	if ns.NodeMap == nil {
		ns.NodeMap = map[string]*Node{}
	}
}

func (ns *Nodes) Node(name string) *Node {
	ns.init()
	n := ns.NodeMap[name]
	if n == nil {
		n = &Node{Name: name}
		ns.NodeMap[name] = n
	}
	return n
}

func main() {
	fmt.Printf("inputs: %v\n", inputs)

	nodes := Nodes{}

	for _, link := range inputs {
		parts := strings.Split(link, ")")
		orbitee := nodes.Node(parts[0])
		orbiter := nodes.Node(parts[1])
		orbitee.Orbits = append(orbitee.Orbits, orbiter)
		orbiter.Orbits = append(orbiter.Orbits, orbitee)
	}

	you := nodes.Node("YOU")
	orbitCount, found := find(you, nil, "SAN")
	fmt.Println(orbitCount-1, found)
}

func find(start *Node, prev *Node, target string) (orbits int, found bool) {
	for _, o := range start.Orbits {
		if o == prev {
			continue
		}
		if o.Name == target {
			return 0, true
		}
		count, found := find(o, start, target)
		if found {
			return 1 + count, found
		}
	}
	return -1, false
}
