package day2

import (
	"testing"
)

func TestDive(t *testing.T) {
	if dive("input1.txt") != 150 {
		t.Fatal()
	}
	if dive("input.txt") != 1855814 {
		t.Fatal()
	}
}

func TestDive_Part2(t *testing.T) {
	if dive_Part2("input1.txt") != 900 {
		t.Fatal()
	}
	if dive_Part2("input.txt") != 1845455714 {
		t.Fatal()
	}
}
