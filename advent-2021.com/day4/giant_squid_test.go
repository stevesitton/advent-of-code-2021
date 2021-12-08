package day4

import (
	"testing"
)

func TestGiantSquid(t *testing.T) {
	if GiantSquid("input1.txt") != 4512 {
		t.Fatal()
	}
	if GiantSquid("input.txt") != 58412 {
		t.Fatal()
	}
}

func TestGiantSquid_Part2(t *testing.T) {
	if GiantSquid_Part2("input1.txt") != 1924 {
		t.Fatal()
	}
	if GiantSquid_Part2("input.txt") != 10030 {
		t.Fatal()
	}
}
