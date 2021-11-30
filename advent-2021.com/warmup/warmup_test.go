package warmup

import (
	"testing"
)

func TestPasswordPhilosopy(t *testing.T) {
	if passwordPhilosophy("passwords1.txt") != 2 {
		t.Fatal()
	}
	if passwordPhilosophy("passwords.txt") != 424 {
		t.Fatal()
	}
}

func TestPasswordPhilosopy_P2(t *testing.T) {
	if passwordPhilosophy_P2("passwords1.txt") != 1 {
		t.Fatal()
	}
	if passwordPhilosophy_P2("passwords.txt") != 747 {
		t.Fatal()
	}
}

func TestReportRepair(t *testing.T) {
	if reportRepair("elf_expenses1.txt") != 514579 {
		t.Fatal()
	}
	if reportRepair("elf_expenses2.txt") != -1 {
		t.Fatal()
	}
	if reportRepair("elf_expenses3.txt") != 1020100 {
		t.Fatal()
	}
	if reportRepair("elf_expenses.txt") != 73371 {
		t.Fatal()
	}
}

func TestReportRepair_P2(t *testing.T) {
	if reportRepair_P2("elf_expenses1.txt") != 241861950 {
		t.Fatal()
	}
	if reportRepair_P2("elf_expenses.txt") != 127642310 {
		t.Fatal()
	}
}

func TestElfCookieCount(t *testing.T) {
	if elfCookieCount("cookies1.txt") != 137 {
		t.Fatal()
	}
	if elfCookieCount("cookies.txt") != 549096 {
		t.Fatal()
	}
}
