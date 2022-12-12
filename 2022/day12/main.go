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
	uniquePaths := make(map[string]struct{})
	pos := newPath(start, len(heightMap[0]), len(heightMap))
	allPaths := pos.Next()
	for {
		steps++
		if len(allPaths) == 0 {
			return 0
		}

		var next []path
		for _, p := range allPaths {
			if _, ok := uniquePaths[p.String()]; ok {
				continue
			}
			uniquePaths[p.String()] = struct{}{}

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

type position struct {
	X, Y int
}

func (c position) Equal(d position) bool {
	return c.X == d.X && c.Y == d.Y
}

func (c position) String() string {
	return fmt.Sprintf("%d/%d", c.X, c.Y)
}

type path struct {
	current    position
	prev       position
	visited    map[string]struct{}
	xLen, yLen int
}

func newPath(c position, xLen, yLen int) path {
	return path{
		current: c,
		prev:    c,
		visited: make(map[string]struct{}),
		xLen:    xLen,
		yLen:    yLen,
	}
}

func (p path) String() string {
	return fmt.Sprintf("%s--%s", p.current, p.prev)
}

func (p path) Clone() path {
	visitedCopy := make(map[string]struct{})
	for k, v := range p.visited {
		visitedCopy[k] = v
	}
	return path{
		current: p.current,
		prev:    p.current,
		visited: visitedCopy,
		xLen:    p.xLen,
		yLen:    p.yLen,
	}
}

func (p path) Next() []path {
	var next []path

	right := p.Clone()
	right.current.X++
	if right.current.X < p.xLen {
		key := right.current.String()
		if _, ok := right.visited[key]; !ok {
			right.visited[key] = struct{}{}
			next = append(next, right)
		}
	}

	left := p.Clone()
	left.current.X--
	if left.current.X >= 0 {
		key := left.current.String()
		if _, ok := left.visited[key]; !ok {
			left.visited[key] = struct{}{}
			next = append(next, left)
		}
	}

	down := p.Clone()
	down.current.Y++
	if down.current.Y < down.yLen {
		key := down.current.String()
		if _, ok := down.visited[key]; !ok {
			down.visited[key] = struct{}{}
			next = append(next, down)
		}
	}

	up := p.Clone()
	up.current.Y--
	if up.current.Y >= 0 {
		key := up.current.String()
		if _, ok := up.visited[key]; !ok {
			up.visited[key] = struct{}{}
			next = append(next, up)
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
				val = []byte("a")[0]
				fallthrough
			case "a":
				start = append(start, position{c, r})
			case "E":
				end = position{c, r}
				val = []byte("z")[0]
			}

			final[r][c] = val
		}
	}
	if !anyStart {
		start = []position{initialStart}
	}

	return final, start, end
}
