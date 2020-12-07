package part1

import (
	"strings"
)

type Batch []TravelDocument

func ParseBatch(input string) Batch {
	batch := Batch{}
	inputs := strings.Split(input, "\n\n")
	for _, i := range inputs {
		td := ParseTravelDocument(i)
		batch = append(batch, td)
	}
	return batch
}

func (b Batch) CountPassports() int {
	valid := 0
	for _, td := range b {
		if td.Passport() {
			valid++
		}
	}
	return valid
}
