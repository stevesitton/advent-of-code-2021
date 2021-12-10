package day5

import (
	"testing"
)

func TestHydrothermalVenture(t *testing.T) {
	if HydrothermalVenture("input1.txt") != 5 {
		t.Fatal()
	}
	if HydrothermalVenture("input.txt") != 5169 {
		t.Fatal()
	}
}

func TestHydrothermalVenture_Part2(t *testing.T) {
	if HydrothermalVenture_Part2("input1.txt") != 12 {
		t.Fatal()
	}
	if HydrothermalVenture_Part2("input.txt") != 22083 {
		t.Fatal()
	}
}
