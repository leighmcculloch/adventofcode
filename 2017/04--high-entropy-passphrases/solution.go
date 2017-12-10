// Package solution contains solutions to the problems described at http://adventofcode.com/2017/day/4.
package solution

import (
	"sort"
	"strings"
)

// Part1.
func Part1(passphraseList string) int {
	passphrases := strings.Split(passphraseList, "\n")
	valid := 0

iteratingPassphrases:
	for _, p := range passphrases {
		words := map[string]bool{}
		for _, w := range strings.Fields(p) {
			if words[w] {
				continue iteratingPassphrases
			}
			words[w] = true
		}
		if len(words) > 0 {
			valid++
		}
	}
	return valid
}

// Part2.
func Part2(passphraseList string) int {
	passphrases := strings.Split(passphraseList, "\n")
	valid := 0

iteratingPassphrases:
	for _, p := range passphrases {
		words := map[string]bool{}
		for _, w := range strings.Fields(p) {
			chars := strings.Split(w, "")
			sort.Strings(chars)
			sortedWord := strings.Join(chars, "")
			if words[sortedWord] {
				continue iteratingPassphrases
			}
			words[sortedWord] = true
		}
		if len(words) > 0 {
			valid++
		}
	}
	return valid
}
