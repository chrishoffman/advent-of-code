package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

type coordinate struct{ X, Y int }

func newCoordinate() *coordinate {
	return &coordinate{0, 0}
}

func (c *coordinate) diff(o *coordinate) coordinate {
	return coordinate{pointDiff(c.X, o.X), pointDiff(c.Y, o.Y)}
}

func (c *coordinate) move(direction string) {
	for _, d := range direction {
		switch string(d) {
		case "U":
			c.Y++
		case "D":
			c.Y--
		case "R":
			c.X++
		case "L":
			c.X--
		}
	}
}

func (c *coordinate) String() string {
	return fmt.Sprintf("%d/%d", c.X, c.Y)
}

func pointDiff(p1, p2 int) int {
	if (p1 <= 0 && p2 <= 0) || (p1 >= 0 && p2 >= 0) {
		return p1 - p2
	}

	direction := 1
	if p1 < p2 {
		direction = -1
	}
	if p1 < 0 {
		p1 *= -1
	}
	if p2 < 0 {
		p2 *= -1
	}
	return (p1 + p2) * direction
}

func problemOne() {
	steps := parseFile()
	hc, tc := newCoordinate(), newCoordinate()

	unique := map[string]bool{}
	unique[tc.String()] = true
	for _, s := range steps {
		for c := 0; c < s.steps; c++ {
			hc.move(s.direction)

			diff := hc.diff(tc)
			if diff.X == 0 {
				if diff.Y > 1 {
					tc.move("U")
				} else if diff.Y < -1 {
					tc.move("D")
				}
			} else if diff.Y == 0 {
				if diff.X > 1 {
					tc.move("R")
				} else if diff.X < -1 {
					tc.move("L")
				}
			} else if diff.X > 1 {
				tc.move("R")
				if diff.Y > 0 {
					tc.move("U")
				} else {
					tc.move("D")
				}
			} else if diff.X < -1 {
				tc.move("L")
				if diff.Y > 0 {
					tc.move("U")
				} else {
					tc.move("D")
				}
			} else if diff.Y > 1 {
				tc.move("U")
				if diff.X > 0 {
					tc.move("R")
				} else {
					tc.move("L")
				}
			} else if diff.Y < -1 {
				tc.move("D")
				if diff.X > 0 {
					tc.move("R")
				} else {
					tc.move("L")
				}
			}

			unique[tc.String()] = true
		}
	}

	fmt.Println(len(unique))
}

func problemTwo() {
}

type step struct {
	direction string
	steps     int
}

func parseFile() []step {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var steps []step
	for _, r := range raw {
		values := strings.Fields(r)
		stepCount, _ := strconv.Atoi(values[1])
		steps = append(steps, step{values[0], stepCount})
	}

	return steps
}
