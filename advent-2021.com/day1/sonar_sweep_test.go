package day1

import (
	"testing"
)

func TestSonarSweep(t *testing.T) {
	if SonarSweep("report1.txt") != 7 {
		t.Fatal()
	}
	if SonarSweep("report.txt") != 1266 {
		t.Fatal()
	}
}

func TestSonarSweep_Part2(t *testing.T) {
	if SonarSweep_Part2("report1.txt") != 5 {
		t.Fatal()
	}
	if SonarSweep_Part2("report.txt") != 1217 {
		t.Fatal()
	}
}
