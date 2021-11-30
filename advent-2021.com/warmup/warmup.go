package warmup

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Password Philosophy
// Your flight departs in a few days from the coastal airport; the easiest way down
// to the coast from here is via toboggan.

// The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day. "Something's
// wrong with our computers; we can't log in!" You ask if you can take a look.

// Their password database seems to be a little corrupted: some of the passwords wouldn't
// have been allowed by the Official Toboggan Corporate Policy that was in effect when
// they were chosen.

// To try to debug the problem, they have created a list (your puzzle input) of passwords
// (according to the corrupted database) and the corporate policy when that password was
// set.

// For example, suppose you have the following list:

// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc
// Each line gives the password policy and then the password. The password policy
// indicates the lowest and highest number of times a given letter must appear for
// the password to be valid. For example, 1-3 a means that the password must contain a
// at least 1 time and at most 3 times.

// In the above example, 2 passwords are valid. The middle password, cdefg, is not; it
// contains no instances of b, but needs at least 1. The first and third passwords are
// valid: they contain one a or nine c, both within the limits of their respective
// policies.

// How many passwords are valid according to their policies?
func passwordPhilosophy(input string) int {

	passwordsFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer passwordsFile.Close()

	scanner := bufio.NewScanner(passwordsFile)
	r, _ := regexp.Compile("[- :]")
	validPasswords := 0
	for scanner.Scan() {
		entry := scanner.Text()
		match := r.Split(entry, 5)
		low, _ := strconv.Atoi(match[0])
		high, _ := strconv.Atoi(match[1])
		letter := match[2]
		password := match[4]
		//fmt.Printf("Low:%v High:%v Letter:%v Password:%v\n", low, high, letter, password)
		count := 0
		for i := 0; i < len(password); i++ {
			if password[i] == letter[0] {
				count++
			}
		}
		if count >= low && count <= high {
			validPasswords++
		}
	}
	//fmt.Println(validPasswords)
	return validPasswords
}

// Password Philosophy - Part 2
// While it appears you validated the passwords correctly, they don't seem to be
// what the Official Toboggan Corporate Authentication System is expecting.

// The shopkeeper suddenly realizes that he just accidentally explained the password
// policy rules from his old job at the sled rental place down the street! The
// Official Toboggan Corporate Policy actually works a little differently.

// Each policy actually describes two positions in the password, where 1 means the
// first character, 2 means the second character, and so on. (Be careful; Toboggan
// Corporate Policies have no concept of "index zero"!) Exactly one of these
// positions must contain the given letter. Other occurrences of the letter are
// irrelevant for the purposes of policy enforcement.

// Given the same example list from above:

// 1-3 a: abcde is valid: position 1 contains a and position 3 does not.
// 1-3 b: cdefg is invalid: neither position 1 nor position 3 contains b.
// 2-9 c: ccccccccc is invalid: both position 2 and position 9 contain c.
// How many passwords are valid according to the new interpretation of the policies?
func passwordPhilosophy_P2(input string) int {

	passwordsFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer passwordsFile.Close()

	scanner := bufio.NewScanner(passwordsFile)
	r, _ := regexp.Compile("[- :]")
	validPasswords := 0
	for scanner.Scan() {
		entry := scanner.Text()
		match := r.Split(entry, 5)
		position1, _ := strconv.Atoi(match[0])
		position2, _ := strconv.Atoi(match[1])
		letter, password := match[2], match[4]
		match1, match2 := false, false

		if string(password[position1-1]) == letter {
			match1 = true
		}
		if string(password[position2-1]) == letter {
			match2 = true
		}
		if (match1 || match2) && match1 != match2 {
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
	return validPasswords
}

// Report Repair
// Specifically, they need you to find the two entries that sum to 2020 and then
// multiply those two numbers together.

// For example, suppose your expense report contained the following:

// 1721
// 979
// 366
// 299
// 675
// 1456
// In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying
// them together produces 1721 * 299 = 514579, so the correct answer is 514579.

// Of course, your expense report is much larger. Find the two entries that sum
// to 2020; what do you get if you multiply them together?
func reportRepair(input string) int {
	expenseFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer expenseFile.Close()

	scanner := bufio.NewScanner(expenseFile)
	var entries []int
	for scanner.Scan() {
		var entry int
		if tmpI, err := strconv.Atoi(scanner.Text()); err == nil {
			entries = append(entries, tmpI)
			entry = tmpI
		}
		// scan map for each entry read in, maybe get match before scanning whole file?
		dupeHit := false
		for _, item := range entries {
			// there will always be one dupe hit for each entry
			if !dupeHit && item == entry {
				dupeHit = true
			} else if item+entry == 2020 {
				fmt.Println(item * entry)
				return item * entry
			}
		}

	}
	return -1
}

// Report Repair - Part 2
// The Elves in accounting are thankful for your help; one of them even offers
// you a starfish coin they had left over from a past vacation. They offer you
// a second one if you can find three numbers in your expense report that meet
// the same criteria.

// Using the above example again, the three entries that sum to 2020 are 979,
// 366, and 675. Multiplying them together produces the answer, 241861950.
func reportRepair_P2(input string) int {
	expenseFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer expenseFile.Close()

	scanner := bufio.NewScanner(expenseFile)
	var entries = make(map[int]int)
	for scanner.Scan() {
		if tmpI, err := strconv.Atoi(scanner.Text()); err == nil {
			entries[tmpI] = tmpI
		}
	}

	// scan map for each entry read in, maybe hit a match before scanning whole file?
	for one := range entries {
		for two := range entries {
			difference := 2020 - (one + two)
			if difference > 0 {
				_, found := entries[difference]
				if found {
					fmt.Println(one * two * difference)
					return one * two * difference
				}
			}
		}
	}
	return -1
}

// It's cookie day at Elf central, and all the elves have written how many
// cookies each of them would like to eat in a spreadsheet.  The head cookie
// baker comes to you with the spreadsheet and asks you to calculate how
// many cookies they'll need to bake in total.  "No problem!" you say, and
// try to load the spreadsheet in excel.  It fails: the elves have terrible
// separator hygiene, and they've put lots of weird column symbols around
// their numbers.  Elves are so weird.  Anyway, you read the file manually
// and realize that all you have to do is add up the first number on every
// line.  How many cookies in total will the elves need?
func elfCookieCount(input string) int {
	cookieFile, err := os.Open(input)
	if err != nil {
		return 0
	}
	defer cookieFile.Close()

	scanner := bufio.NewScanner(cookieFile)
	count := 0
	r, _ := regexp.Compile("[0-9]+")
	for scanner.Scan() {
		match := r.FindStringSubmatch(scanner.Text())
		if i, err := strconv.Atoi(match[0]); err == nil {
			count += i
		}
	}

	fmt.Println("Total: ", count)
	return count
}
