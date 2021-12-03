package day3

import (
	"testing"
)

func TestBinaryDiagnostics(t *testing.T) {
	if BinaryDiagnostics("input1.txt") != 198 {
		t.Fatal()
	}
	if BinaryDiagnostics("input.txt") != 1082324 {
		t.Fatal()
	}
}

func TestBinaryDiagnostics_Part2(t *testing.T) {
	if BinaryDiagnostics_Part2("input1.txt") != 230 {
		t.Fatal()
	}
	if BinaryDiagnostics_Part2("input.txt") != 1353024 {
		t.Fatal()
	}
}
