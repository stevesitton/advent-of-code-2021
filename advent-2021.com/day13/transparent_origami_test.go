package day13

import (
	"testing"
)

func TestTransparentOrigami(t *testing.T) {
	if TransparentOrigami("input_example.txt") != 17 {
		t.Fatal()
	}
	if TransparentOrigami("input.txt") != 621 {
		t.Fatal()
	}
}

func TestTransparentOrigami_Part2(t *testing.T) {
	if TransparentOrigami_Part2("input_example.txt") != 16 {
		t.Fatal()
	}
	if TransparentOrigami_Part2("input.txt") != 95 {
		t.Fatal()
	} //HKUJGAJZ
}
