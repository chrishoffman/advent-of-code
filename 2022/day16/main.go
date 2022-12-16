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

	currentValve := "AA"
	maxTime := 30

	var total int
	for i := maxTime; i > 0; i-- {
		fmt.Println("------------------------", currentValve)

		var maxValue, pathLen int
		closedValves := volcano.closedValves()
		if len(closedValves) == 0 {
			fmt.Println(total)
			return
		}
		for c := 0; c < len(closedValves); c++ {
			cost := volcano.shortestPathCost(currentValve, closedValves[c], i)
			value := (i - cost) * volcano[closedValves[c]].flow
			if value > maxValue {
				maxValue = value
				pathLen = cost
				currentValve = closedValves[c]
			}
			fmt.Println(closedValves[c], cost, value, i, volcano[closedValves[c]].flow)
		}
		i += pathLen
		total += maxValue
		volcano[currentValve].state = open
	}

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

func (v volcano) shortestPathCost(valve1, valve2 string, max int) int {
	var cost int
	visted := make(map[string]struct{})

	next := v[valve1].tunnels
	for {
		cost++
		if cost == max {
			return max
		}

		for _, t := range next {
			if t == valve2 {
				return cost
			}

			if _, ok := visted[t]; ok {
				continue
			}

			next = append(next, v[t].tunnels...)
		}
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
