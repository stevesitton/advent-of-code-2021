package day11

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DumboOctopus(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	energyGrid := make(map[int][]int)
	scanner := bufio.NewScanner(inputFile)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := []int{}
		for _, num := range strings.Split(line, "") {
			j, _ := strconv.Atoi(num)
			numbers = append(numbers, j)
		}
		energyGrid[counter] = numbers
		counter++
	}

	steps, totalFlashes := 0, 0
	for {
		if steps == 100 {
			break
		}

		flashes := make(map[string]int)
		for i := 0; i < len(energyGrid); i++ {
			for j := 0; j < len(energyGrid[i]); j++ {
				if energyGrid[i][j] == 9 {
					flashes[fmt.Sprint([]int{i, j})] = 1 // can be any value here, the key is key!
					energyGrid[i][j] = 0
					CheckSurroundsValues([]int{i, j}, energyGrid, flashes)
				} else if _, ok := flashes[fmt.Sprint([]int{i, j})]; !ok {
					energyGrid[i][j]++
				}
			}
		}
		totalFlashes += len(flashes)
		steps++
	}

	//fmt.Println(totalFlashes)
	return totalFlashes
}

func DumboOctopus_Part2(input string) int {
	inputFile, err := os.Open(input)
	if err != nil {
		return -1
	}
	defer inputFile.Close()

	energyGrid := make(map[int][]int)
	scanner := bufio.NewScanner(inputFile)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := []int{}
		for _, num := range strings.Split(line, "") {
			j, _ := strconv.Atoi(num)
			numbers = append(numbers, j)
		}
		energyGrid[counter] = numbers
		counter++
	}

	steps := 0
	for {
		flashes := make(map[string]int)
		for i := 0; i < len(energyGrid); i++ {
			for j := 0; j < len(energyGrid[i]); j++ {
				if energyGrid[i][j] == 9 {
					flashes[fmt.Sprint([]int{i, j})] = 1 // can be any value here, the key is key!
					energyGrid[i][j] = 0
					CheckSurroundsValues([]int{i, j}, energyGrid, flashes)
				} else if _, ok := flashes[fmt.Sprint([]int{i, j})]; !ok {
					energyGrid[i][j]++
				}
			}
		}
		if len(flashes) == 100 {
			break
		}
		steps++
	}
	//fmt.Println(steps + 1)
	return steps + 1
}

func CheckSurroundsValues(start []int, grid map[int][]int, flashes map[string]int) {
	var pointsToCheck Stack
	pointsToCheck.Push(start)
	for {
		if pointsToCheck.Len() == 0 {
			break
		}

		p, _ := pointsToCheck.Pop() // pointer
		if p[0]-1 >= 0 {
			if p[1]-1 >= 0 {
				CheckEnergyValue([]int{p[0] - 1, p[1] - 1}, grid, flashes, &pointsToCheck)
			}
			CheckEnergyValue([]int{p[0] - 1, p[1]}, grid, flashes, &pointsToCheck)
			if p[1]+1 < len(grid[0]) {
				CheckEnergyValue([]int{p[0] - 1, p[1] + 1}, grid, flashes, &pointsToCheck)
			}
		}

		if p[1]-1 >= 0 {
			CheckEnergyValue([]int{p[0], p[1] - 1}, grid, flashes, &pointsToCheck)
		}
		if p[1]+1 < len(grid[0]) {
			CheckEnergyValue([]int{p[0], p[1] + 1}, grid, flashes, &pointsToCheck)
		}

		if p[0]+1 < len(grid) {
			if p[1]-1 >= 0 {
				CheckEnergyValue([]int{p[0] + 1, p[1] - 1}, grid, flashes, &pointsToCheck)
			}
			CheckEnergyValue([]int{p[0] + 1, p[1]}, grid, flashes, &pointsToCheck)
			if p[1]+1 < len(grid[0]) {
				CheckEnergyValue([]int{p[0] + 1, p[1] + 1}, grid, flashes, &pointsToCheck)
			}
		}
	}
}

func CheckEnergyValue(pointer []int, grid map[int][]int, flashes map[string]int,
	pointsToCheck *Stack) {
	key := fmt.Sprint(pointer)
	if _, ok := flashes[key]; !ok {
		if grid[pointer[0]][pointer[1]] == 9 {
			flashes[key] = 1 // can be any value here, the key is key!
			grid[pointer[0]][pointer[1]] = 0
			pointsToCheck.Push(pointer)
		} else {
			grid[pointer[0]][pointer[1]]++
		}
	}
}

// Stack implementation
type Stack [][]int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(point []int) {
	*s = append(*s, point)
}

func (s *Stack) Pop() ([]int, bool) {
	if s.Len() == 0 {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
