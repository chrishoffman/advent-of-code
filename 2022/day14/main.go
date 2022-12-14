package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

type material int

const (
	air material = iota
	rock
	sand
)

func problemOne() {
	grid := populateGrid(false)
	units := fill(grid)
	fmt.Println(units)
}

func problemTwo() {
	grid := populateGrid(true)
	units := fill(grid)
	fmt.Println(units)
}

func fill(grid [][]material) int {
	var units int
ABYSS:
	for {
		current := coordinate{500, 0}

		var moves int
		for {
			var moved bool
			grid[current.Y][current.X] = sand
			for _, m := range []int{0, -1, 1} {
				next := coordinate{current.X + m, current.Y + 1}

				if grid[next.Y][next.X] == air {
					grid[current.Y][current.X] = air
					current = next
					moved = true
					moves++

					if current.Y >= len(grid)-1 {
						break ABYSS
					}
					break
				}
			}

			if !moved {
				break
			}
		}

		units++
		if moves == 0 {
			break
		}
	}
	if len(os.Args) > 1 {
		printGrid(grid)
	}

	return units
}

func populateGrid(addFloor bool) [][]material {
	structures, maxX, maxY := parseFile()

	if addFloor {
		maxX = 2 * maxX
		maxY += 2
		structures = append(structures, structure([]coordinate{{0, maxY}, {maxX, maxY}}))
	}

	grid := make([][]material, maxY+1)
	for r := 0; r < len(grid); r++ {
		// adding 2 because we are evaluating next move left and right
		grid[r] = make([]material, maxX+2)
	}

	for _, s := range structures {
		for l := 0; l < len(s)-1; l++ {
			xPath := []int{s[l].X, s[l+1].X}
			sort.Ints(xPath)
			yPath := []int{s[l].Y, s[l+1].Y}
			sort.Ints(yPath)
			for x := xPath[0]; x <= xPath[1]; x++ {
				for y := yPath[0]; y <= yPath[1]; y++ {
					grid[y][x] = rock
				}
			}
		}
	}

	if len(os.Args) > 1 {
		printGrid(grid)
	}

	return grid
}

func printGrid(grid [][]material) {
	for _, r := range grid {
		for c := 0; c < len(r); c++ {
			switch r[c] {
			case air:
				fmt.Print(".")
			case rock:
				fmt.Print("#")
			case sand:
				fmt.Print("o")
			}
		}
		fmt.Print("\n\n")
	}
}

type coordinate struct{ X, Y int }
type structure []coordinate

func parseFile() ([]structure, int, int) {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var structs []structure
	var maxX, maxY int
	for _, r := range raw {
		var cs []coordinate
		coordinates := strings.Split(r, " -> ")
		for _, c := range coordinates {
			pts := strings.Split(c, ",")
			x, _ := strconv.Atoi(pts[0])
			y, _ := strconv.Atoi(pts[1])
			cs = append(cs, coordinate{x, y})

			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
		structs = append(structs, cs)
	}

	return structs, maxX, maxY
}
