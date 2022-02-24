package day15

import (
	"testing"
)

func TestChiton(t *testing.T) {
	if Chiton("input_example.txt") != 40 {
		t.Fatal()
	}
	if Chiton("input.txt") != 540 {
		t.Fatal()
	}
}
