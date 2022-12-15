package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	cnt, _, _ := scan(2000000, math.MinInt, math.MaxInt)
	fmt.Println(cnt)
}

func problemTwo() {
	// sensors := parseFile()
	// max := 20

	// for y := 0; y < max+1; y++ {
	// NEXTX:
	// 	for x := 0; x < max+1; x++ {
	// 		// find closest sensor
	// 		var closest *sensor
	// 		minDist := math.MaxInt
	// 		for _, s := range sensors {
	// 			if s.pt1.X == x && s.pt1.Y == y {
	// 				break NEXTX
	// 			}

	// 			dist := sensor{coordinate{x, y}, s.pt1}.distance()
	// 			if dist < minDist {
	// 				minDist = dist
	// 				closest = &s
	// 			}
	// 		}

	// 		if closest == nil {
	// 			fmt.Println(x*4000000 + y)
	// 			return
	// 		}

	// 		yDiff := intDiff(closest.pt1.Y, y)
	// 		xDiff := minDist - yDiff

	// 		if xDiff > 0 {
	// 			x += xDiff * 2
	// 		}
	// 	}
	// }
}

func scan(y, minBound, maxBound int) (int, int, map[int]struct{}) {
	sensors := parseFile()
	beaconsOnPath := make(map[int]struct{})
	for _, s := range sensors {
		if s.pt2.Y == y {
			beaconsOnPath[s.pt2.X] = struct{}{}
		}
	}
	uniqueXCoor := make(map[int]struct{})
	for _, s := range sensors {
		dist := s.distance()
		yDiff := intDiff(s.pt1.Y, y)
		xDiff := dist - yDiff
		for d := 0; d < xDiff+1; d++ {
			for _, p := range []int{s.pt1.X + d, s.pt1.X - d} {
				if p < minBound || p > maxBound {
					continue
				}
				if _, ok := beaconsOnPath[p]; !ok {
					uniqueXCoor[p] = struct{}{}
				}
			}
		}
	}

	return len(uniqueXCoor), len(beaconsOnPath), uniqueXCoor
}

type coordinate struct{ X, Y int }
type sensor struct{ pt1, pt2 coordinate }

func (s sensor) distance() int {
	return intDiff(s.pt1.X, s.pt2.X) + intDiff(s.pt1.Y, s.pt2.Y)
}

func parseFile() []sensor {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var sensors []sensor
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for _, r := range raw {
		m := re.FindStringSubmatch(r)
		sensors = append(sensors, sensor{coordinate{strToInt(m[1]), strToInt(m[2])}, coordinate{strToInt(m[3]), strToInt(m[4])}})
	}
	return sensors
}

func strToInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func intDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
