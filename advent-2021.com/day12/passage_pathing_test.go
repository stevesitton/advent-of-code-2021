package day12

import (
	"testing"
)

func TestPassagePathing(t *testing.T) {
	if PassagePathing("input_example.txt") != 10 {
		t.Fatal()
	}
	if PassagePathing("input_example2.txt") != 19 {
		t.Fatal()
	}
	if PassagePathing("input_example3.txt") != 226 {
		t.Fatal()
	}
	if PassagePathing("input.txt") != 3887 {
		t.Fatal()
	}
}

func TestPassagePathing_Part2(t *testing.T) {
	if PassagePathing_Part2("input_example.txt") != 36 {
		t.Fatal()
	}
	if PassagePathing_Part2("input_example2.txt") != 103 {
		t.Fatal()
	}
	if PassagePathing_Part2("input_example3.txt") != 3509 {
		t.Fatal()
	}
	if PassagePathing_Part2("input.txt") != 104834 {
		t.Fatal()
	}
}
