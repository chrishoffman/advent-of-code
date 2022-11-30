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
		var xIncrement, yIncrement int
		switch {
		case line.p1.x > line.p2.x:
			xIncrement = -1
		case line.p1.x < line.p2.x:
			xIncrement = 1
		}

		switch {
		case line.p1.y > line.p2.y:
			yIncrement = -1
		case line.p1.y < line.p2.y:
			yIncrement = 1
		}

		var x, y int = line.p1.x, line.p1.y
		for !(x == line.p2.x+xIncrement && y == line.p2.y+yIncrement) {
			plot[x][y]++
			x += xIncrement
			y += yIncrement
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
