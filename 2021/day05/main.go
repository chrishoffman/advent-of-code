package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

type coordinate struct{ x, y int }
type line struct{ p1, p2 coordinate }

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	lines := parseFile()

	var plot [1000][1000]int

	for _, line := range lines {
		switch {
		case line.p1.x == line.p2.x:
			minCoordinate := min(line.p1.y, line.p2.y)
			maxCoordinate := max(line.p1.y, line.p2.y)
			for y := minCoordinate; y <= maxCoordinate; y++ {
				plot[line.p1.x][y]++
			}
		case line.p1.y == line.p2.y:
			minCoordinate := min(line.p1.x, line.p2.x)
			maxCoordinate := max(line.p1.x, line.p2.x)
			for x := minCoordinate; x <= maxCoordinate; x++ {
				plot[x][line.p1.y]++
			}
		}
	}

	var count int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if plot[x][y] > 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func problemTwo() {
	lines := parseFile()

	var plot [1000][1000]int

	for _, line := range lines {
		switch {
		case line.p1.x == line.p2.x:
			ys := []int{line.p1.y, line.p2.y}
			sort.Ints(ys)
			for y := ys[0]; y <= ys[1]; y++ {
				plot[line.p1.x][y]++
			}
		case line.p1.y == line.p2.y:
			xs := []int{line.p1.x, line.p2.x}
			sort.Ints(xs)
			for x := xs[0]; x <= xs[1]; x++ {
				plot[x][line.p1.y]++
			}
		default:
			var xIncrement, yIncrement int = 1, 1
			if line.p1.x > line.p2.x {
				xIncrement = -1
			}
			if line.p1.y > line.p2.y {
				yIncrement = -1
			}

			var currX, currY int = line.p1.x, line.p1.y
			for currX != line.p2.x+xIncrement && currY != line.p2.y+yIncrement {
				plot[currX][currY]++
				currX += xIncrement
				currY += yIncrement
			}

		}
	}

	var count int
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if plot[x][y] > 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func parseFile() []line {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var lines []line
	for _, l := range raw {
		coordinate1 := parseCoordinate(strings.Split(l, " -> ")[0])
		coordinate2 := parseCoordinate(strings.Split(l, " -> ")[1])

		lines = append(lines, line{coordinate1, coordinate2})
	}

	return lines
}

func parseCoordinate(raw string) coordinate {
	splitCoordinate := strings.Split(raw, ",")
	xRaw, yRaw := splitCoordinate[0], splitCoordinate[1]

	x, err := strconv.Atoi(xRaw)
	if err != nil {
		log.Fatalln(err)
	}

	y, err := strconv.Atoi(yRaw)
	if err != nil {
		log.Fatalln(err)
	}

	return coordinate{x, y}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
