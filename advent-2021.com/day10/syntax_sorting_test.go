package day10

import (
	"testing"
)

func TestSyntaxSorting(t *testing.T) {
	if SyntaxSorting("input1.txt") != 26397 {
		t.Fatal()
	}
	if SyntaxSorting("input.txt") != 392139 {
		t.Fatal()
	}
}

func TestSyntaxSorting_Part2(t *testing.T) {
	if SyntaxSorting_Part2("input1.txt") != 288957 {
		t.Fatal()
	}
	if SyntaxSorting_Part2("input.txt") != 4001832844 {
		t.Fatal()
	}
}
