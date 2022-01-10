package day11

import (
	"testing"
)

func TestDumboOctopus(t *testing.T) {
	if DumboOctopus("input1.txt") != 1656 {
		t.Fatal()
	}
	if DumboOctopus("input.txt") != 1741 {
		t.Fatal()
	}
}

func TestDumboOctopus_Part2(t *testing.T) {
	if DumboOctopus_Part2("input1.txt") != 195 {
		t.Fatal()
	}
	if DumboOctopus_Part2("input.txt") != 440 {
		t.Fatal()
	}
}
