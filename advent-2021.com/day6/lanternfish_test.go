package day6

import (
	"testing"
)

func TestLanternfish(t *testing.T) {
	if Lanternfish("input1.txt", 18) != 26 {
		t.Fatal()
	}
	if Lanternfish("input1.txt", 80) != 5934 {
		t.Fatal()
	}
	if Lanternfish("input.txt", 80) != 374994 {
		t.Fatal()
	}
}

func TestLanternfish_Part2(t *testing.T) {
	if Lanternfish_Part2("input1.txt", 18) != 26 {
		t.Fatal()
	}
	if Lanternfish_Part2("input1.txt", 80) != 5934 {
		t.Fatal()
	}
	if Lanternfish_Part2("input1.txt", 256) != 26984457539 {
		t.Fatal()
	}
	if Lanternfish_Part2("input.txt", 256) != 1686252324092 {
		t.Fatal()
	}
}
