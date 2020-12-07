package part1

import (
	"io/ioutil"
	"testing"
)

func TestInput(t *testing.T) {
	input, err := ioutil.ReadFile("testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	batch := ParseBatch(string(input))
	passports := batch.CountPassports()
	wantPassports := 2
	if passports != wantPassports {
		t.Fatalf("got %v want %v", passports, wantPassports)
	}
}

func TestDonna(t *testing.T) {
	input, err := ioutil.ReadFile("testdata/donna.txt")
	if err != nil {
		t.Fatal(err)
	}

	batch := ParseBatch(string(input))
	passports := batch.CountPassports()
	wantPassports := -1
	if passports != wantPassports {
		t.Fatalf("got %v want %v", passports, wantPassports)
	}
}

func TestLeigh(t *testing.T) {
	input, err := ioutil.ReadFile("testdata/leigh.txt")
	if err != nil {
		t.Fatal(err)
	}

	batch := ParseBatch(string(input))
	passports := batch.CountPassports()
	wantPassports := -1
	if passports != wantPassports {
		t.Fatalf("got %v want %v", passports, wantPassports)
	}
}
