package day17

import (
	"testing"
)

func TestTrickShot(t *testing.T) {
	// if TrickShot("target area: x=20..30, y=-10..-5") != 45 {
	// 	t.Fatal()
	// }
	if TrickShot("target area: x=88..125, y=-157..-103") != 12246 {
		t.Fatal()
	}
}

func TestTrickShot_Part2(t *testing.T) {
	// if TrickShot_Part2("target area: x=20..30, y=-10..-5") != 112 {
	// 	t.Fatal()
	// }
	if TrickShot_Part2("target area: x=88..125, y=-157..-103") != 3528 {
		t.Fatal()
	}
}
