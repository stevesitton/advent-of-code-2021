package day9

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
--- Day 9: Smoke Basin ---
These caves seem to be lava tubes. Parts are even still volcanically active; small hydrothermal
vents release smoke into the caves that slowly settles like rain.

If you can model how the smoke flows through the caves, you might be able to avoid it and be
that much safer. The submarine generates a heightmap of the floor of the nearby caves for you
(your puzzle input).

Smoke flows to the lowest point of the area it's in. For example, consider the following
heightmap:

2199943210
3987894921
9856789892
8767896789
9899965678

Each number corresponds to the height of a particular location, where 9 is the highest and 0 is
the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than any of its adjacent
locations. Most locations have four adjacent locations (up, down, left, and right); locations on
the edge or corner of the map have three or two adjacent locations, respectively. (Diagonal
locations do not count as adjacent.)

In the above example, there are four low points, all highlighted: two are in the first row (a 1
	and a 0), one is in the third row (a 5), and one is in the bottom row (also a 5). All other
	locations on the heightmap have some lower adjacent location, and so are not low points.

The risk level of a low point is 1 plus its height. In the above example, the risk levels of the
low points are 2, 1, 6, and 6. The sum of the risk levels of all low points in the heightmap is
therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk levels of all low
points on your heightmap? */
func SmokeBasin(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	heightmap := make(map[int][]int)
	scanner := bufio.NewScanner(inputFile)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := []int{}
		for _, num := range strings.Split(line, "") {
			j, _ := strconv.Atoi(num)
			numbers = append(numbers, j)
		}
		heightmap[counter] = numbers
		counter++
	}

	pointer := []int{0, 0}
	output := 0
	for {
		if pointer[0] == len(heightmap) {
			break
		}

		height := heightmap[pointer[0]][pointer[1]]
		surroundingHeights := make([]int, 0)
		if pointer[0]-1 >= 0 {
			if pointer[1]-1 >= 0 {
				surroundingHeights = append(surroundingHeights, heightmap[pointer[0]-1][pointer[1]-1])
			}
			surroundingHeights = append(surroundingHeights, heightmap[pointer[0]-1][pointer[1]])
			if pointer[1]+1 < len(heightmap[0]) {
				surroundingHeights = append(surroundingHeights, heightmap[pointer[0]-1][pointer[1]+1])
			}
		}

		if pointer[1]-1 >= 0 {
			surroundingHeights = append(surroundingHeights, heightmap[pointer[0]][pointer[1]-1])
		}
		if pointer[1]+1 < len(heightmap[0]) {
			surroundingHeights = append(surroundingHeights, heightmap[pointer[0]][pointer[1]+1])
		}

		if pointer[0]+1 < len(heightmap) {
			if pointer[1]-1 >= 0 {
				surroundingHeights = append(surroundingHeights, heightmap[pointer[0]+1][pointer[1]-1])
			}
			surroundingHeights = append(surroundingHeights, heightmap[pointer[0]+1][pointer[1]])
			if pointer[1]+1 < len(heightmap[0]) {
				surroundingHeights = append(surroundingHeights, heightmap[pointer[0]+1][pointer[1]+1])
			}
		}

		isLowest := true
		for _, h := range surroundingHeights {
			if h <= height {
				isLowest = false
				break
			}
		}
		if isLowest {
			output += 1 + height
		}

		pointer[1]++
		if pointer[1] == len(heightmap[0]) {
			if pointer[0] < len(heightmap) {
				pointer[0]++
				pointer[1] = 0
			}
		}
	}
	//fmt.Println(output)
	return output
}

/*
--- Part Two ---
Next, you need to find the largest basins so you know what areas are most important to avoid.

A basin is all locations that eventually flow downward to a single low point. Therefore, every
 low point has a basin, although some basins are very small. Locations of height 9 do not count
  as being in any basin, and all other locations will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the low point. The
example above has four basins.

The top-left basin, size 3:

2199943210
3987894921
9856789892
8767896789
9899965678
The top-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
The middle basin, size 14:

2199943210
3987894921
9856789892
8767896789
9899965678
The bottom-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
Find the three largest basins and multiply their sizes together. In the above example,
this is 9 * 14 * 9 = 1134.

What do you get if you multiply together the sizes of the three largest basins? */
func SmokeBasin_Part2(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	heightmap := make(map[int][]int)
	scanner := bufio.NewScanner(inputFile)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := []int{}
		for _, num := range strings.Split(line, "") {
			j, _ := strconv.Atoi(num)
			numbers = append(numbers, j)
		}
		heightmap[counter] = numbers
		counter++
	}

	lowestPoints := GetLowestPoints(heightmap)

	basinSizes := GetBasinSizes(lowestPoints, heightmap)
	values := []int{}
	for _, value := range basinSizes {
		values = append(values, value)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	return values[0] * values[1] * values[2]
}

func GetBasinSizes(lowestPoints [][]int, heightmap map[int][]int) map[int]int {
	basinSizes := make(map[int]int)
	for i, lowestPoint := range lowestPoints {
		basinSizes[i] = GetBasinSize(lowestPoint, heightmap)
	}
	return basinSizes
}

func GetBasinSize(lowest []int, heightmap map[int][]int) int {

	basinHeights := make(map[string]int)
	pointsToCheck := make([][]int, 1)
	pointsToCheck[0] = lowest
	basinHeights[fmt.Sprint(lowest)] = heightmap[lowest[0]][lowest[1]]
	for {
		if len(pointsToCheck) == 0 {
			break
		}

		pointer := pointsToCheck[0]
		height := heightmap[pointer[0]][pointer[1]]

		if pointer[0]-1 >= 0 {
			adjPointer := []int{pointer[0] - 1, pointer[1]}
			if CheckAdjacentHeight(heightmap[pointer[0]-1][pointer[1]], height, adjPointer, basinHeights, pointsToCheck) {
				pointsToCheck = append(pointsToCheck, adjPointer)
			}
		}

		if pointer[1]-1 >= 0 {
			adjPointer := []int{pointer[0], pointer[1] - 1}
			if CheckAdjacentHeight(heightmap[pointer[0]][pointer[1]-1], height, adjPointer, basinHeights, pointsToCheck) {
				pointsToCheck = append(pointsToCheck, adjPointer)
			}
		}
		if pointer[1]+1 < len(heightmap[0]) {
			adjPointer := []int{pointer[0], pointer[1] + 1}
			if CheckAdjacentHeight(heightmap[pointer[0]][pointer[1]+1], height, adjPointer, basinHeights, pointsToCheck) {
				pointsToCheck = append(pointsToCheck, adjPointer)
			}
		}

		if pointer[0]+1 < len(heightmap) {
			adjPointer := []int{pointer[0] + 1, pointer[1]}
			if CheckAdjacentHeight(heightmap[pointer[0]+1][pointer[1]], height, adjPointer, basinHeights, pointsToCheck) {
				pointsToCheck = append(pointsToCheck, adjPointer)
			}
		}

		if len(pointsToCheck) == 1 {
			pointsToCheck = nil
		} else {
			pointsToCheck = pointsToCheck[1:]
		}
	}
	return len(basinHeights)
}

func CheckAdjacentHeight(adjHeight int, height int, adjPointer []int,
	basinHeights map[string]int, pointsToCheck [][]int) bool {

	if adjHeight != 9 {
		if _, found := basinHeights[fmt.Sprint(adjPointer)]; !found {
			basinHeights[fmt.Sprint(adjPointer)] = adjHeight
			return true
		}
	}
	return false
}

func GetLowestPoints(heightmap map[int][]int) [][]int {

	pointer := []int{0, 0}
	lowestPoints := make([][]int, 0)
	output := 0
	for {
		if pointer[0] == len(heightmap) {
			break
		}

		height := heightmap[pointer[0]][pointer[1]]
		surroundingHeights := make([]int, 0)
		if pointer[0]-1 >= 0 {
			if pointer[1]-1 >= 0 {
				surroundingHeights = append(surroundingHeights, heightmap[pointer[0]-1][pointer[1]-1])
			}
			surroundingHeights = append(surroundingHeights, heightmap[pointer[0]-1][pointer[1]])
			if pointer[1]+1 < len(heightmap[0]) {
				surroundingHeights = append(surroundingHeights, heightmap[pointer[0]-1][pointer[1]+1])
			}
		}

		if pointer[1]-1 >= 0 {
			surroundingHeights = append(surroundingHeights, heightmap[pointer[0]][pointer[1]-1])
		}
		if pointer[1]+1 < len(heightmap[0]) {
			surroundingHeights = append(surroundingHeights, heightmap[pointer[0]][pointer[1]+1])
		}

		if pointer[0]+1 < len(heightmap) {
			if pointer[1]-1 >= 0 {
				surroundingHeights = append(surroundingHeights, heightmap[pointer[0]+1][pointer[1]-1])
			}
			surroundingHeights = append(surroundingHeights, heightmap[pointer[0]+1][pointer[1]])
			if pointer[1]+1 < len(heightmap[0]) {
				surroundingHeights = append(surroundingHeights, heightmap[pointer[0]+1][pointer[1]+1])
			}
		}

		isLowest := true
		for _, h := range surroundingHeights {
			if h <= height {
				isLowest = false
				break
			}
		}

		if isLowest {
			pointerCopy := []int{0, 0}
			copy(pointerCopy, pointer)
			lowestPoints = append(lowestPoints, pointerCopy)
			output += 1 + height
		}

		pointer[1]++
		if pointer[1] == len(heightmap[0]) {
			if pointer[0] < len(heightmap) {
				pointer[0]++
				pointer[1] = 0
			}
		}
	}
	return lowestPoints
}
