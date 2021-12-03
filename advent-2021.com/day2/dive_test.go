package day2

import (
	"testing"
)

func TestDive(t *testing.T) {
	if Dive("input1.txt") != 150 {
		t.Fatal()
	}
	if Dive("input.txt") != 1855814 {
		t.Fatal()
	}
}

func TestDive_Part2(t *testing.T) {
	if Dive_Part2("input1.txt") != 900 {
		t.Fatal()
	}
	if Dive_Part2("input.txt") != 1845455714 {
		t.Fatal()
	}
}
