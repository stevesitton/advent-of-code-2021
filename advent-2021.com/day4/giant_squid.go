package day4

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
--- Day 4: Giant Squid ---
You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that
you can't see any sunlight. What you can see, however, is a giant squid that has attached itself
to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen
at random, and the chosen number is marked on all boards on which it appears. (Numbers may not
	appear on all boards.) If all numbers in any row or any column of a board are marked, that
	board wins. (Diagonals don't count.)

The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass
the time. It automatically generates a random order in which to draw numbers and a random set of
boards (your puzzle input). For example:

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards
are marked as follows (shown here adjacent to each other to save space):

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
Finally, 24 is drawn:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
At this point, the third board wins because it has at least one complete row or column of marked
numbers (in this case, the entire top row is marked: 14 21 17 24 4).

The score of the winning board can now be calculated. Start by finding the sum of all unmarked
numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that
was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win first. What will
your final score be if you choose that board?
*/
func GiantSquid(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	body, _ := ioutil.ReadFile(input)
	r, _ := regexp.Compile("\n\n")
	r2, _ := regexp.Compile(",")
	bodySplit := r.Split(string(body), 30)

	bingoNumbers := r2.Split(bodySplit[0], 100)
	boards := GenerateBoards(bodySplit)

	for _, numberStr := range bingoNumbers {
		//fmt.Println("Number:", numberStr)
		for j, board := range boards {
			hit := false
			for k := 0; k < len(board); k++ {
				row := board[k]
				for l, numberNum := range row {
					if numberNum == numberStr {
						hit = true
						row[l] = "*"
						break
					}
				}
			}
			if hit && IsBingoBoard(boards[j]) {
				sum := CountUnmarkedNumbers(boards[j])
				number, _ := strconv.Atoi(numberStr)
				fmt.Println(number)
				fmt.Println(sum * number)
				return sum * number
			}
		}
	}
	return -1
}

/*
--- Part Two ---
On the other hand, it might be wise to try a different strategy: let the giant squid win.

You aren't sure how many bingo boards a giant squid could play at once, so rather than waste
time counting its arms, the safe thing to do is to figure out which board will win last and
choose that one. That way, no matter which boards it picks, it will win for sure.

In the above example, the second board is the last to win, which happens after 13 is
eventually called and its middle column is completely marked. If you were to keep playing until
this point, the second board would have a sum of unmarked numbers equal to 148 for a final
score of 148 * 13 = 1924.

Figure out which board will win last. Once it wins, what would its final score be?
*/
func GiantSquid_Part2(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	body, _ := ioutil.ReadFile(input)
	r, _ := regexp.Compile("\n\n")
	r2, _ := regexp.Compile(",")
	bodySplit := r.Split(string(body), 50)

	bingoNumbers := r2.Split(bodySplit[0], 100)
	boards := GenerateBoards(bodySplit)

	for _, numberStr := range bingoNumbers {
		for j, board := range boards {
			hit := false
			for k := 0; k < len(board); k++ {
				row := board[k]
				for l, numberNum := range row {
					if numberNum == numberStr {
						hit = true
						row[l] = "*"
						break
					}
				}
			}
			if hit && IsBingoBoard(board) {
				if len(boards) == 1 {
					for _, winner := range boards {
						// there will only be one!
						sum := CountUnmarkedNumbers(winner)
						number, _ := strconv.Atoi(numberStr)
						// fmt.Println(sum * number)
						return sum * number
					}
				}
				// remove board
				delete(boards, j)
			}
		}
	}
	return -1
}

func CountUnmarkedNumbers(board map[int][]string) int {
	fmt.Println("WINNING BOARD...")
	DisplayBoard(board)
	sum := 0
	for _, row := range board {
		for _, numberNum := range row {
			if numberNum != "*" {
				number, _ := strconv.Atoi(numberNum)
				sum += number
			}
		}
	}
	return sum
}

func DisplayBoard(board map[int][]string) {
	for _, row := range board {
		fmt.Println(row)
	}
}

func IsBingoBoard(board map[int][]string) bool {
	// check rows
	starCount := 0
	for _, row := range board {
		for _, numberNum := range row {
			if numberNum != "*" {
				break
			} else {
				starCount++
			}
		}
		if starCount == 5 {
			return true
		}
		starCount = 0
	}

	// check columns
	for i := 0; i < 5; i++ {
		for _, row := range board {
			if row[i] != "*" {
				break
			} else {
				starCount++
			}
		}
		if starCount == 5 {
			return true
		}
		starCount = 0
	}

	return false
}

func GenerateBoards(input []string) map[int]map[int][]string {
	r2, _ := regexp.Compile("\n")
	r3, _ := regexp.Compile("[ ]+")
	boards := make(map[int]map[int][]string)
	for i := 1; i < len(input); i++ {
		boardStr := input[i]
		rows := r2.Split(string(boardStr), 5)
		board := make(map[int][]string)
		for j, row := range rows {
			board[j] = r3.Split(strings.TrimSpace(row), 5)
		}
		boards[i-1] = board
	}
	return boards
}
