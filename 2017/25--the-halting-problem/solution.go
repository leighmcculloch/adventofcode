// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/25.
package solution

func Solution() int {
	t := tape{}
	s := state(a{})
	i := 0

	for stepCount := 0; stepCount < 12134527; stepCount++ {
		var newValue bool
		var move int

		if t.Get(i) {
			newValue, move, s = s.DoForTrue()
		} else {
			newValue, move, s = s.DoForFalse()
		}

		t.Set(i, newValue)
		i += move
	}

	return t.Checksum()
}

type tape map[int]bool

func (t tape) Set(i int, b bool) {
	if b {
		t[i] = true
	} else {
		delete(t, i)
	}
}

func (t tape) Get(i int) bool {
	return t[i]
}

func (t tape) Checksum() int {
	return len(t)
}

type state interface {
	DoForFalse() (newValue bool, move int, newState state)
	DoForTrue() (newValue bool, move int, newState state)
}

type a struct{}

func (a) DoForFalse() (newValue bool, move int, newState state) { return true, +1, b{} }
func (a) DoForTrue() (newValue bool, move int, newState state)  { return false, -1, c{} }

type b struct{}

func (b) DoForFalse() (newValue bool, move int, newState state) { return true, -1, a{} }
func (b) DoForTrue() (newValue bool, move int, newState state)  { return true, +1, c{} }

type c struct{}

func (c) DoForFalse() (newValue bool, move int, newState state) { return true, +1, a{} }
func (c) DoForTrue() (newValue bool, move int, newState state)  { return false, -1, d{} }

type d struct{}

func (d) DoForFalse() (newValue bool, move int, newState state) { return true, -1, e{} }
func (d) DoForTrue() (newValue bool, move int, newState state)  { return true, -1, c{} }

type e struct{}

func (e) DoForFalse() (newValue bool, move int, newState state) { return true, +1, f{} }
func (e) DoForTrue() (newValue bool, move int, newState state)  { return true, +1, a{} }

type f struct{}

func (f) DoForFalse() (newValue bool, move int, newState state) { return true, +1, a{} }
func (f) DoForTrue() (newValue bool, move int, newState state)  { return true, +1, e{} }
