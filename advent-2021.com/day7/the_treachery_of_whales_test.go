package day7

import (
	"testing"
)

func TestTheTreacheryOfWhales(t *testing.T) {
	if TheTreacheryOfWhales("input1.txt") != 37 {
		t.Fatal()
	}
	if TheTreacheryOfWhales("input.txt") != 357353 {
		t.Fatal()
	}
}

func TestTheTreacheryOfWhales_Part2(t *testing.T) {
	if TheTreacheryOfWhales_Part2("input1.txt") != 168 {
		t.Fatal()
	}
	if TheTreacheryOfWhales_Part2("input.txt") != 104822130 {
		t.Fatal()
	}
}
