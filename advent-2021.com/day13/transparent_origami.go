package day13

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func TransparentOrigami(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	dots := make(map[int][]int)
	folds := make([]string, 0)
	r, _ := regexp.Compile(",")
	r2, _ := regexp.Compile("=")
	counter, highestX, highestY := 0, 0, 0
	for scanner.Scan() {
		entry := scanner.Text()
		if len(entry) != 0 && !strings.Contains(entry, "fold") {
			coord := r.Split(entry, 3)
			coordX, _ := strconv.Atoi(coord[0])
			coordY, _ := strconv.Atoi(coord[1])
			dots[counter] = []int{coordX, coordY}

			if coordX > highestX {
				highestX = coordX
			}
			if coordY > highestY {
				highestY = coordY
			}
		} else if len(entry) != 0 && strings.Contains(entry, "fold") {
			fold := r2.Split(entry, 2)
			folds = append(folds, fold[0][len(fold[0])-1:]+"="+fold[1])
			break
		}
		counter++
	}

	grid := BuildGrid(highestX, highestY)
	AddDotsToGrid(grid, dots)
	for _, nextFold := range folds {
		fold := r2.Split(nextFold, 3)
		foldX := true
		if fold[0] == "y" {
			foldX = false
		}
		foldLine, _ := strconv.Atoi(fold[1])
		if !foldX {
			rowCounter := 0
			for i := len(grid) - 1; i >= foldLine; i-- {
				if i != foldLine {
					for j := 0; j < len(grid[i]); j++ {
						if grid[i][j] == 1 {
							grid[rowCounter][j] = 1
						}
					}
				}
				rowCounter++
				delete(grid, i)
			}
		} else {
			colCounter := 0
			for i := len(grid[0]) - 1; i >= foldLine; i-- {
				if i != foldLine {
					for j := 0; j < len(grid); j++ {
						if grid[j][i] == 1 {
							grid[j][colCounter] = 1
						}
					}
				}
				colCounter++
			}
			// more efficient to remove cols with one loop than with each iteration above
			for i := 0; i < len(grid); i++ {
				grid[i] = grid[i][:foldLine]
			}
		}
	}

	fmt.Println(CountGrid(grid))
	return CountGrid(grid)
}

// Needed to change the way I folded for Part 2!
func TransparentOrigami_Part2(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	dots := make(map[int][]int)
	folds := make([]string, 0)
	r, _ := regexp.Compile(",")
	r2, _ := regexp.Compile("=")
	counter, counterFolds, highestX, highestY := 0, 0, 0, 0
	for scanner.Scan() {
		entry := scanner.Text()
		if len(entry) != 0 && !strings.Contains(entry, "fold") {
			coord := r.Split(entry, 3)
			coordX, _ := strconv.Atoi(coord[0])
			coordY, _ := strconv.Atoi(coord[1])
			dots[counter] = []int{coordX, coordY}

			if coordX > highestX {
				highestX = coordX
			}
			if coordY > highestY {
				highestY = coordY
			}
		} else if len(entry) != 0 && strings.Contains(entry, "fold") {
			fold := r2.Split(entry, 2)
			folds = append(folds, fold[0][len(fold[0])-1:]+"="+fold[1])
			counterFolds++
		}
		counter++
	}

	grid := BuildGrid(highestX, highestY)
	AddDotsToGrid(grid, dots)

	for _, nextFold := range folds {
		fold := r2.Split(nextFold, 3)
		foldX := true
		if fold[0] == "y" {
			foldX = false
		}
		foldLine, _ := strconv.Atoi(fold[1])
		if !foldX {
			rowCounter := 1
			for i := foldLine + 1; i < len(grid); i++ {
				for j := 0; j < len(grid[i]); j++ {
					if grid[i][j] == 1 {
						grid[foldLine-rowCounter][j] = 1
					}
				}
				rowCounter++
			}
			for i := len(grid); i >= foldLine; i-- {
				delete(grid, i)
			}
		} else {
			colCounter := 1
			for i := foldLine + 1; i < len(grid[0]); i++ {
				for j := 0; j < len(grid); j++ {
					if grid[j][i] == 1 {
						grid[j][foldLine-colCounter] = 1
					}
				}
				colCounter++
			}
			// more efficient to remove cols with one loop than with each iteration above
			for i := 0; i < len(grid); i++ {
				grid[i] = grid[i][:foldLine]
			}
		}
	}
	fmt.Printf("========\n")
	PrintGrid(grid)

	fmt.Println(CountGrid(grid))
	return CountGrid(grid)
}

func AddDotsToGrid(grid map[int][]int, vents map[int][]int) {
	for _, vent := range vents {
		grid[vent[1]][vent[0]] = 1
	}
}

func BuildGrid(highX int, highY int) map[int][]int {
	grid := make(map[int][]int)
	for i := 0; i <= highY; i++ {
		row := make([]int, highX+1)
		grid[i] = row
	}
	return grid
}

func CountGrid(grid map[int][]int) int {
	gridCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				gridCount++
			}
		}
	}
	return gridCount
}

func PrintGrid(grid map[int][]int) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
}
