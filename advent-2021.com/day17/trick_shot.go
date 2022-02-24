package day17

import (
	"fmt"
	"strconv"
	"strings"
)

func TrickShot(input string) int {

	type pointer struct{ x, y int }

	xStr := input[strings.Index(input, "=")+1 : strings.Index(input, ",")]
	x1, _ := strconv.Atoi(xStr[:strings.Index(xStr, ".")])
	x2, _ := strconv.Atoi(xStr[strings.Index(xStr, ".")+2:])
	targetX := pointer{x1, x2}

	yStr := input[strings.Index(input, "y=")+2:]
	y1, _ := strconv.Atoi(yStr[:strings.Index(yStr, ".")])
	y2, _ := strconv.Atoi(yStr[strings.Index(yStr, ".")+2:])
	targetY := pointer{y1, y2}

	fireProbe := func(v pointer) int {
		maxY := 0
		p := pointer{v.x, v.y}
		isHit := false
		for {
			if v.x != 0 {
				if v.x > 0 {
					v.x--
				} else {
					v.x++
				}
			}
			v.y--
			p.x += v.x
			p.y += v.y
			if p.y > maxY {
				maxY = p.y
			}
			//fmt.Println(p)

			// if hit target
			if p.x >= targetX.x && p.x <= targetX.y &&
				p.y >= targetY.x && p.y <= targetY.y {
				//fmt.Println("HIT")
				isHit = true
				break
			}

			// if missed target
			if p.x > targetX.y || p.y < targetY.y {
				//fmt.Println("MISS")
				break
			}

		}
		if !isHit {
			maxY = -1
		}
		//fmt.Println(maxY)
		return maxY
	}

	maxY := 0
	for x := 2; x < 500; x++ {
		for y := 1; y < 300; y++ {
			//fmt.Println("Fire: ", x, y)
			tmpY := fireProbe(pointer{x, y})
			if tmpY > maxY {
				maxY = tmpY
			}
		}
	}

	fmt.Println(maxY)
	return maxY
}

func TrickShot_Part2(input string) int {

	type pointer struct{ x, y int }

	xStr := input[strings.Index(input, "=")+1 : strings.Index(input, ",")]
	x1, _ := strconv.Atoi(xStr[:strings.Index(xStr, ".")])
	x2, _ := strconv.Atoi(xStr[strings.Index(xStr, ".")+2:])
	targetX := pointer{x1, x2}

	yStr := input[strings.Index(input, "y=")+2:]
	y1, _ := strconv.Atoi(yStr[:strings.Index(yStr, ".")])
	y2, _ := strconv.Atoi(yStr[strings.Index(yStr, ".")+2:])
	targetY := pointer{y1, y2}

	fireProbe := func(v pointer) bool {
		missCount := 0
		p := pointer{v.x, v.y}
		for {
			// if hit target
			if p.x >= targetX.x && p.x <= targetX.y &&
				p.y >= targetY.x && p.y <= targetY.y {
				return true
			}

			// if missed target
			if p.x > targetX.y || p.y < targetY.y {
				missCount++
				if missCount == 3 {
					break
				}
			}

			if v.x != 0 {
				if v.x > 0 {
					v.x--
				} else {
					v.x++
				}
			}
			v.y--
			p.x += v.x
			p.y += v.y
		}
		return false
	}

	hitCount := 0
	for x := 2; x < 500; x++ {
		for y := -200; y < 400; y++ {
			if fireProbe(pointer{x, y}) {
				hitCount++
			}
		}
	}

	//fmt.Println(hitCount)
	return hitCount
}
