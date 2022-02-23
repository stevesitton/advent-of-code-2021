package day14

import (
	"testing"
)

func TestExtendedPolymerization(t *testing.T) {
	if ExtendedPolymerization("input_example.txt", 10) != 1588 {
		t.Fatal()
	}
	if ExtendedPolymerization("input.txt", 10) != 3213 {
		t.Fatal()
	}
}

func TestExtendedPolymerization_Part2(t *testing.T) {
	if ExtendedPolymerization_Part2("input_example.txt", 10) != 1588 {
		t.Fatal()
	}
	if ExtendedPolymerization_Part2("input.txt", 10) != 3213 {
		t.Fatal()
	}
	if ExtendedPolymerization_Part2("input_example.txt", 40) != 2188189693529 {
		t.Fatal()
	}
	if ExtendedPolymerization_Part2("input.txt", 40) != 3711743744429 {
		t.Fatal()
	}
}
