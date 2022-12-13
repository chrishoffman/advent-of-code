package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	heightMap, starts, end := parseFile(false)
	fmt.Println(findShortestPath(heightMap, starts[0], end))
}

func problemTwo() {
	heightMap, starts, end := parseFile(true)
	fmt.Println(findMinShortestPath(heightMap, starts, end))
}

func findMinShortestPath(heightMap [][]byte, starts []position, end position) int {
	var dists []int
	for _, start := range starts {
		dist := findShortestPath(heightMap, start, end)
		if dist > 0 {
			dists = append(dists, dist)
		}
	}
	sort.Ints(dists)
	return dists[0]
}

func findShortestPath(heightMap [][]byte, start, end position) int {
	var steps int
	visited := make(map[string]struct{})
	pos := path{
		current: start,
		xLen:    len(heightMap[0]),
		yLen:    len(heightMap),
	}
	allPaths := pos.Next()

	for {
		steps++
		// deadend path
		if len(allPaths) == 0 {
			return 0
		}

		var next []path
		for _, p := range allPaths {
			if _, ok := visited[p.String()]; ok {
				continue
			}
			visited[p.String()] = struct{}{}

			currHeight := heightMap[p.current.Y][p.current.X]
			prevHeight := heightMap[p.prev.Y][p.prev.X]
			if int(currHeight)-int(prevHeight) < 2 {
				if p.current.Equal(end) {
					return steps
				}
				next = append(next, p.Next()...)
			}
		}
		allPaths = next
	}
}

type position struct{ X, Y int }

func (c position) Equal(d position) bool {
	return c.X == d.X && c.Y == d.Y
}

func (c position) String() string {
	return fmt.Sprintf("%d/%d", c.X, c.Y)
}

type path struct {
	current, prev position
	xLen, yLen    int
}

func (p path) String() string {
	return fmt.Sprintf("%s--%s", p.current, p.prev)
}

func (p path) Next() []path {
	var next []path

	dirs := []position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range dirs {
		new := position{p.current.X + dir.X, p.current.Y + dir.Y}
		if new.X >= 0 && new.X < p.xLen && new.Y >= 0 && new.Y < p.yLen {
			next = append(next, path{
				current: new,
				prev:    p.current,
				xLen:    p.xLen,
				yLen:    p.yLen,
			})
		}
	}

	return next
}

func parseFile(anyStart bool) ([][]byte, []position, position) {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	final := make([][]byte, len(raw))
	var start []position
	var initialStart, end position
	for r := 0; r < len(raw); r++ {
		final[r] = make([]byte, len(raw[r]))
		for c := 0; c < len(final[r]); c++ {
			val := raw[r][c]
			switch string(val) {
			case "S":
				initialStart = position{c, r}
				val = byte('a')
				fallthrough
			case "a":
				start = append(start, position{c, r})
			case "E":
				end = position{c, r}
				val = byte('z')
			}
			final[r][c] = val
		}
	}
	if !anyStart {
		start = []position{initialStart}
	}

	return final, start, end
}
