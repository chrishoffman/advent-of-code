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
	volcano.relievePressure()
	fmt.Println(volcano.totalFlow)
}

func problemTwo() {
}

type volcano struct {
	current    string
	valves     map[string]*valve
	openValves int
	flow       int
	totalFlow  int
	timeLeft   int
}

func (v *volcano) relievePressure() {
	v.current = "AA"
	v.timeLeft = 30
	v.next()
}

func (v *volcano) next() {
	if v.timeLeft <= 0 {
		return
	}

	v.timeLeft--
	v.totalFlow += v.flow

	if v.openValves == len(v.valves) {
		v.next()
	}

	fmt.Println(v.timeLeft, v.openValves, len(v.valves))

	var candidate []*volcano
	if v.valves[v.current].state == closed {
		volcano := v.clone()
		volcano.timeLeft--
		volcano.valves[volcano.current].open()
		volcano.openValves++
		volcano.flow += volcano.valves[volcano.current].flow
		volcano.next()
		candidate = append(candidate, volcano)
	}
	for _, t := range v.valves[v.current].tunnels {
		if v.valves[t].state == open {
			continue
		}
		volcano := v.clone()
		volcano.current = t
		volcano.next()
		candidate = append(candidate, volcano)
	}

	max := &volcano{}
	for _, c := range candidate {
		if c.totalFlow > max.totalFlow {
			v.totalFlow = max.totalFlow
			v.valves = max.valves
			v.flow = max.flow
			v.timeLeft = max.timeLeft
			fmt.Println(c.totalFlow)
		}
	}
}

func (v *volcano) clone() *volcano {
	new := make(map[string]*valve)
	for k, v := range v.valves {
		new[k] = v.clone()
	}
	return &volcano{
		valves:     new,
		current:    v.current,
		flow:       v.flow,
		openValves: v.openValves,
		totalFlow:  v.totalFlow,
		timeLeft:   v.timeLeft,
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

func (v *valve) open() {
	v.state = open
}

func (v valve) clone() *valve {
	c := &valve{
		flow:  v.flow,
		state: v.state,
	}
	c.tunnels = append(c.tunnels, v.tunnels...)
	return c
}

func parseFile() *volcano {
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
	return &volcano{
		valves: valves,
	}
}
