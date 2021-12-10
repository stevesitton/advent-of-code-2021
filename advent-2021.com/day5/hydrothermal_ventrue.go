package day5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

/*
--- Day 5: Hydrothermal Venture ---
You come across a field of hydrothermal vents on the ocean floor! These vents constantly
produce large, opaque clouds, so it would be best to avoid them if possible.

They tend to form in lines; the submarine helpfully produces a list of nearby lines of
vents (your puzzle input) for you to review. For example:

0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2

Each line of vents is given as a line segment in the format x1,y1 -> x2,y2 where x1,y1 are
the coordinates of one end the line segment and x2,y2 are the coordinates of the other end.
These line segments include the points at both ends. In other words:

An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.
For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.

So, the horizontal and vertical lines from the above list would produce the following
diagram:

.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....

In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9. Each position
is shown as the number of lines which cover that point or . if no line covers that point.
The top-left pair of 1s, for example, comes from 2,2 -> 2,1; the very bottom row is formed
by the overlapping lines 0,9 -> 5,9 and 0,9 -> 2,9.

To avoid the most dangerous areas, you need to determine the number of points where at least
two lines overlap. In the above example, this is anywhere in the diagram with a 2 or larger
 - a total of 5 points.

Consider only horizontal and vertical lines. At how many points do at least two lines
overlap? */
func HydrothermalVenture(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	vents := make(map[int][]int)
	r, _ := regexp.Compile(" -> ")
	r2, _ := regexp.Compile(",")
	counter, highestX, highestY := 0, 0, 0
	for scanner.Scan() {
		entry := scanner.Text()
		coord := r.Split(entry, 3)
		start := r2.Split(coord[0], 2)
		end := r2.Split(coord[1], 2)
		fromX, _ := strconv.Atoi(start[0])
		fromY, _ := strconv.Atoi(start[1])
		toX, _ := strconv.Atoi(end[0])
		toY, _ := strconv.Atoi(end[1])
		if fromX == toX || fromY == toY {
			// non-diagonal vent
			if fromX == toX && fromY > toY {
				tmp := fromY
				fromY = toY
				toY = tmp
			} else if fromY == toY && fromX > toX {
				tmp := fromX
				fromX = toX
				toX = tmp
			}
			if toX > highestX {
				highestX = toX
			}
			if toY > highestY {
				highestY = toY
			}
			//fmt.Printf("Start:%v %v End:%v %v\n", fromX, fromY, toX, toY)
			vents[counter] = []int{fromX, fromY, toX, toY}
			counter++
		}
	}

	grid := BuildGrid(highestX, highestY)
	AddVentsToGrid(grid, vents)

	ventOverlapCount := GetVentOverlapCount(grid)
	//fmt.Println(ventOverlapCount)
	return ventOverlapCount
}

/*
--- Part Two ---
Unfortunately, considering only horizontal and vertical lines doesn't give you the full
picture; you need to also consider diagonal lines.

Because of the limits of the hydrothermal vent mapping system, the lines in your list
will only ever be horizontal, vertical, or a diagonal line at exactly 45 degrees. In
other words:

An entry like 1,1 -> 3,3 covers points 1,1, 2,2, and 3,3.
An entry like 9,7 -> 7,9 covers points 9,7, 8,8, and 7,9.
Considering all lines from the above example would now produce the following diagram:

1.1....11.
.111...2..
..2.1.111.
...1.2.2..
.112313211
...1.2....
..1...1...
.1.....1..
1.......1.
222111....
You still need to determine the number of points where at least two lines overlap. In
the above example, this is still anywhere in the diagram with a 2 or larger - now a
total of 12 points.

Consider all of the lines. At how many points do at least two lines overlap? */
func HydrothermalVenture_Part2(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	vents := make(map[int][]int)
	r, _ := regexp.Compile(" -> ")
	r2, _ := regexp.Compile(",")
	counter, highestX, highestY := 0, 0, 0
	for scanner.Scan() {
		entry := scanner.Text()
		coord := r.Split(entry, 3)
		start := r2.Split(coord[0], 2)
		end := r2.Split(coord[1], 2)
		fromX, _ := strconv.Atoi(start[0])
		fromY, _ := strconv.Atoi(start[1])
		toX, _ := strconv.Atoi(end[0])
		toY, _ := strconv.Atoi(end[1])
		if fromX == toX || fromY == toY {
			if fromX == toX && fromY > toY {
				tmp := fromY
				fromY = toY
				toY = tmp
			} else if fromY == toY && fromX > toX {
				tmp := fromX
				fromX = toX
				toX = tmp
			}
		}
		if toX > highestX {
			highestX = toX
		}
		if toY > highestY {
			highestY = toY
		}
		vents[counter] = []int{fromX, fromY, toX, toY}
		counter++
	}

	grid := BuildGrid(highestX, highestY)
	AddVentsToGrid(grid, vents)

	ventOverlapCount := GetVentOverlapCount(grid)
	//fmt.Println(ventOverlapCount)
	return ventOverlapCount
}

func AddVentsToGrid(grid map[int][]int, vents map[int][]int) {
	// add vents to grid
	for _, vent := range vents {
		// 0:fromX  1:fromY  2:toX  3:toY
		pointerX := vent[0]
		pointerY := vent[1]
		for {
			grid[pointerY][pointerX]++
			if pointerX == vent[2] && pointerY == vent[3] {
				break
			}
			if pointerX < vent[2] {
				pointerX++
			} else if pointerX > vent[2] {
				pointerX--
			}

			if pointerY < vent[3] {
				pointerY++
			} else if pointerY > vent[3] {
				pointerY--
			}
		}
	}
}

func GetVentOverlapCount(grid map[int][]int) int {
	count := 0
	for _, row := range grid {
		for _, vent := range row {
			if vent > 1 {
				count++
			}
		}
	}
	return count
}

func BuildGrid(highX int, highY int) map[int][]int {
	grid := make(map[int][]int)
	for i := 0; i <= highY; i++ {
		row := make([]int, highX+1)
		grid[i] = row
	}
	return grid
}

func PrintGrid(grid map[int][]int) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
}
