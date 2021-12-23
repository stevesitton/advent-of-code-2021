package day9

import (
	"testing"
)

func TestSmokeBasin(t *testing.T) {
	if SmokeBasin("input1.txt") != 15 {
		t.Fatal()
	}
	if SmokeBasin("input.txt") != 554 {
		t.Fatal()
	}
}

func TestSmokeBasin_Part2(t *testing.T) {
	if SmokeBasin_Part2("input1.txt") != 1134 {
		t.Fatal()
	}
	if SmokeBasin_Part2("input.txt") != 1017792 {
		t.Fatal()
	}
}
