// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/16.
package solution

import (
	"bytes"
	"fmt"
	"strings"
)

func Part1(starting, dance string) string {
	moves := []move{}
	for _, s := range strings.Split(dance, ",") {
		switch s[0] {
		case 's':
			moves = append(moves, parseSpin(s[1:]))
		case 'x':
			moves = append(moves, parseExchange(s[1:]))
		case 'p':
			moves = append(moves, parsePartner(s[1:]))
		}
	}

	input := []byte(starting)

	for _, m := range moves {
		m.Apply(input)
	}

	return string(input)
}

func Part2(starting, dance string) string {
	moves := []move{}
	for _, s := range strings.Split(dance, ",") {
		switch s[0] {
		case 's':
			moves = append(moves, parseSpin(s[1:]))
		case 'x':
			moves = append(moves, parseExchange(s[1:]))
		case 'p':
			moves = append(moves, parsePartner(s[1:]))
		}
	}

	input := []byte(starting)

	seen := newList()

	const iters = 1000000000
	var i = 0
	for ; i < iters; i++ {
		for _, m := range moves {
			m.Apply(input)
		}

		inputStr := string(input)
		if seen.Contains(inputStr) {
			break
		}
		seen.Add(inputStr)
	}

	return seen.Get(iters%seen.Len() - 1)
}

type move interface {
	Apply(input []byte)
}

type spin struct {
	count int
}

func parseSpin(s string) spin {
	m := spin{}
	fmt.Sscanf(s, "%d", &m.count)
	return m
}

func (m spin) Apply(input []byte) {
	end := input[len(input)-m.count:]
	start := input[:len(input)-m.count]
	newInput := make([]byte, 0, len(input))
	newInput = append(newInput, end...)
	newInput = append(newInput, start...)
	copy(input, newInput)
}

type exchange struct {
	a int
	b int
}

func parseExchange(s string) exchange {
	e := exchange{}
	fmt.Sscanf(s, "%d/%d", &e.a, &e.b)
	return e
}

func (e exchange) Apply(input []byte) {
	input[e.a], input[e.b] = input[e.b], input[e.a]
}

type partner struct {
	a byte
	b byte
}

func parsePartner(s string) partner {
	p := partner{}
	ab := strings.Split(s, "/")
	p.a = ab[0][0]
	p.b = ab[1][0]
	return p
}

func (p partner) Apply(input []byte) {
	a := bytes.IndexByte(input, p.a)
	b := bytes.IndexByte(input, p.b)
	input[a], input[b] = input[b], input[a]
}

type list struct {
	m map[string]bool
	a []string
}

func newList() list {
	return list{m: map[string]bool{}, a: []string{}}
}

func (l *list) Add(s string) {
	l.m[s] = true
	l.a = append(l.a, s)
}

func (l list) Contains(s string) bool {
	return l.m[s]
}

func (l list) Get(i int) string {
	return l.a[i]
}

func (l list) Len() int {
	return len(l.a)
}
