package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	volcano := parseFile()

	for {
		currentValve := "AA"
		maxTime := 30
		totalValves := len(volcano)

		var total int
		for i := maxTime; i > 0; i-- {
			closedValves := volcano.closedValves()
			if len(closedValves) == 0 {
				break
			}

			var next string
			var maxValue, pathLen int
			for c := 0; c < len(closedValves); c++ {
				cost := volcano.shortestPath(currentValve, closedValves[c], i)
				value := (i - cost - 1) * volcano[closedValves[c]].flow
			}
			i -= pathLen
			total += maxValue
			currentValve = next
			volcano[currentValve].state = open
		}
	}
	fmt.Println(total)
}

func problemTwo() {
}

type volcano map[string]*valve

func (v volcano) closedValves() []string {
	var closedValves []string
	for k, valve := range v {
		if valve.state == closed && valve.flow > 0 {
			closedValves = append(closedValves, k)
		}
	}
	return closedValves
}

func (v volcano) shortestPath(valve1, valve2 string, max int) int {
	var cost int
	visited := make(map[string]struct{})

	all := v[valve1].tunnels
	for {
		cost++
		if cost == max {
			return 0
		}

		var next []string
		for _, t := range all {
			if t == valve2 {
				return cost
			}

			if _, ok := visited[t]; ok {
				continue
			}

			visited[t] = struct{}{}
			next = append(next, v[t].tunnels...)
		}
		all = next
	}
}

type valveState uint

const (
	closed valveState = iota
	open
)

type valve struct {
	flow    int
	state   valveState
	tunnels []string
}

func parseFile() volcano {
	raw, _ := advent.ParseFile("./sinput.txt")
	valves := make(map[string]*valve)
	for _, r := range raw {
		words := strings.Fields(strings.Replace(r, ",", "", -1))
		flow, _ := strconv.Atoi(words[4][5 : len(words[4])-1])
		valves[words[1]] = &valve{
			flow:    flow,
			tunnels: words[9:],
		}
	}
	return valves
}
