package day8

import (
	"testing"
)

func TestSevenSegmentSearch(t *testing.T) {
	if SevenSegmentSearch("input1.txt") != 26 {
		t.Fatal()
	}
	if SevenSegmentSearch("input.txt") != 390 {
		t.Fatal()
	}
}

func TestSevenSegmentSearch_Part2(t *testing.T) {
	if SevenSegmentSearch_Part2("input1.txt") != 61229 {
		t.Fatal()
	}
	if SevenSegmentSearch_Part2("input.txt") != 1011785 {
		t.Fatal()
	}
}
