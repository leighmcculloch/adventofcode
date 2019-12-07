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
	Name       string
	Orbits     *Node
	OrbittedBy []*Node
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
		orbiter.Orbits = orbitee
		orbitee.OrbittedBy = append(orbitee.OrbittedBy, orbiter)
	}

	com := nodes.Node("COM")
	orbitCount := countOrbits(com, 0)
	fmt.Println(orbitCount)
}

func countOrbits(n *Node, depth int) int {
	count := 0
	for _, o := range n.OrbittedBy {
		count += 1 + depth + countOrbits(o, depth+1)
	}
	return count
}
