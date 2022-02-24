package day15

import (
	"bytes"
	"fmt"
	"os"
)

/*
Didn't get this to work
*/
func Chiton(input string) int {
	data, _ := os.ReadFile(input)
	grid := bytes.Fields(data)
	fmt.Println(len(grid))

	n := len(grid)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1e9
		}
	}

	type point struct{ x, y int }
	work := make([][]point, 20*n)

	add := func(p point, d int) {
		if p.x < 0 || p.y < 0 || p.x >= n || p.y >= n {
			return
		}
		d += int(grid[p.x][p.y]) - '0'
		if dist[p.x][p.y] <= d {
			return
		}
		dist[p.x][p.y] = d
		work[d] = append(work[d], p)
	}
	add(point{0, 0}, -int(grid[0][0]-'0'))

	visit := func(p point) {
		d := dist[p.x][p.y]
		if p.x == n-1 && p.y == n-1 {
			fmt.Println(d)
			//return d - int(grid[0][0])
			os.Exit(0)
		}
		add(point{p.x - 1, p.y}, d)
		add(point{p.x + 1, p.y}, d)
		add(point{p.x, p.y - 1}, d)
		add(point{p.x, p.y + 1}, d)
	}

	for _, w := range work {
		for _, p := range w {
			visit(p)
		}
	}

	return -1
	/*
		inputFile, err := os.Open(input)
		if err != nil {
			return -1
		}
		defer inputFile.Close()

		scanner := bufio.NewScanner(inputFile)

		grid := make(map[int][]int)
		rowCounter := 0
		for scanner.Scan() {
			entry := scanner.Text()
			row := make([]int, len(entry))
			grid[rowCounter] = row

			for i := 0; i < len(entry); i++ {
				value, _ := strconv.Atoi(string(entry[i]))
				grid[rowCounter][i] = value
			}
			rowCounter++
		}

		distance := make([][]int, len(grid))
		for i := range distance {
			distance[i] = make([]int, len(grid))
			for j := range distance[i] {
				distance[i][j] = 1e7
			}
		}

		type pointer struct{ x, y int }
		queue := make([][]pointer, 20*len(grid))

		add := func(p pointer, d int) {
			if p.x < 0 || p.y < 0 || p.x >= len(grid) || p.y >= len(grid) {
				return
			}
			d += grid[p.x][p.y]
			if distance[p.x][p.y] <= d {
				return
			}
			distance[p.x][p.y] = d
			queue[d] = append(queue[d], p)
		}
		add(pointer{0, 0}, 0)

		for _, q := range queue {
			for _, p := range q {
				//result = CheckPoint(grid, distance, queue, p)
				//fmt.Println(p)
				d := distance[p.x][p.y]
				if p.x == len(grid)-1 && p.y == len(grid)-1 {
					fmt.Println(d - grid[0][0])
					return d - grid[0][0]
				}
				add(pointer{p.x - 1, p.y}, d)
				add(pointer{p.x + 1, p.y}, d)
				add(pointer{p.x, p.y - 1}, d)
				add(pointer{p.x, p.y + 1}, d)
			}
		}

		return -1
	*/
}

// func CheckPoint(grid map[int][]int, distance [][]int, queue [][]pointer, p pointer) int {
// 	d := distance[p.x][p.y]
// 	if p.x == 9 && p.y == 9 {
// 		return d
// 	}
// 	AddToQueue(grid, distance, queue, pointer{p.x - 1, p.y}, d)
// 	AddToQueue(grid, distance, queue, pointer{p.x + 1, p.y}, d)
// 	AddToQueue(grid, distance, queue, pointer{p.x, p.y - 1}, d)
// 	AddToQueue(grid, distance, queue, pointer{p.x, p.y + 1}, d)
// 	return -1
// }

// func AddToQueue(grid map[int][]int, distance [][]int, queue [][]pointer, p pointer, d int) {
// 	if p.x < 0 || p.y < 0 || p.x >= 9 || p.y >= 9 {
// 		return
// 	}
// 	d += grid[p.x][p.y]
// 	if distance[p.x][p.y] <= d {
// 		return
// 	}
// 	distance[p.x][p.y] = d
// 	queue[d] = append(queue[d], p)
// }

//PrintGrid(grid)

//route := make(map[int][]int)

//stepCounter := 0
// for i := 0; i < len(grid); i++ {

// 	for j := 0; j < len(grid[0]); j++ {
// 		if i == 0 && j == 0 {
// 			continue
// 		}
// 		//fmt.Printf("%d,%d: %d\n", i, j, grid[i][j])
// 		if i == 0 {
// 			grid[i][j] += grid[i][j-1]
// 		} else if j == 0 {
// 			grid[i][j] += grid[i-1][j]
// 		} else {
// 			valX := grid[i-1][j]
// 			valY := grid[i][j-1]
// 			if valX < valY {
// 				grid[i][j] += grid[i-1][j]
// 			} else {
// 				grid[i][j] += grid[i][j-1]
// 			}
// 		}
// 	}
// }
// for {
// 	fmt.Println(pointer, grid[pointer[0]][pointer[1]])
// 	nextSquare := FindNextSquare(grid, pointer)
// 	route[stepCounter] = nextSquare
// 	pointer = nextSquare
// 	stepCounter++

// 	if pointer[0] == len(grid)-1 && pointer[1] == len(grid[0])-1 {
// 		break
// 	}
// }

//PrintGrid(grid)
// routeTotal := 0
// for i := 0; i < len(route); i++ {
// 	fmt.Println(route[i])
// 	routeTotal += grid[route[i][0]][route[i][1]]
// }

// 	fmt.Println(grid[len(grid)-1][len(grid[0])-1] - grid[0][0])
// 	return grid[len(grid)-1][len(grid[0])-1] - grid[0][0]
// }

func RecurseSquares(grid map[int][]int, pointer []int) int {
	if pointer[0] == len(grid)-1 && pointer[1] == len(grid[0])-1 {
		fmt.Println("END")
	} else {
		value := grid[pointer[0]][pointer[1]]

		RecurseSquares(grid, []int{pointer[0] + 1, pointer[1]})
		RecurseSquares(grid, []int{pointer[0], pointer[1] + 1})
		return value
	}
	return 0
}

func FindNextSquare(grid map[int][]int, pointer []int) []int {

	if pointer[0]+1 >= len(grid) {
		return []int{pointer[0], pointer[1] + 1}
	} else if pointer[1]+1 >= len(grid[0]) {
		return []int{pointer[0] + 1, pointer[1]}
	}
	below := grid[pointer[0]+1][pointer[1]]
	below2b := grid[pointer[0]+2][pointer[1]]
	below2r := grid[pointer[0]+1][pointer[1]+1]
	b := 0
	if below+below2b <= below+below2r {
		b = below + below2b
	} else {
		b = below + below2r
	}
	fmt.Printf("b:%d below:%d below2b:%d below2r:%d\n", b, below, below2b, below2r)
	right := grid[pointer[0]][pointer[1]+1]
	right2b := grid[pointer[0]+1][pointer[1]+1]
	right2r := grid[pointer[0]][pointer[1]+2]
	r := 0
	if right+right2b <= right+right2r {
		r = right + right2b
	} else {
		r = right + right2r
	}
	fmt.Printf("r:%d right:%d right2b:%d right2r:%d\n", r, right, right2b, right2r)
	if b <= r {
		return []int{pointer[0] + 1, pointer[1]}
	}
	return []int{pointer[0], pointer[1] + 1}
}

func PrintGrid(grid map[int][]int) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
}
