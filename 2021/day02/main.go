package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

type move struct {
	direction string
	steps     int
}

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	values := parseFile()

	var h, d int

	for _, move := range values {
		switch move.direction {
		case "forward":
			h += move.steps
		case "down":
			d += move.steps
		case "up":
			d -= move.steps
		default:
			log.Fatalln("unknown direction")
		}
	}

	fmt.Println(h * d)
}

func problemTwo() {
	values := parseFile()

	var h, d, aim int

	for _, move := range values {
		switch move.direction {
		case "forward":
			h += move.steps
			d += aim * move.steps
		case "down":
			aim += move.steps
		case "up":
			aim -= move.steps
		default:
			log.Fatalln("unknown direction")
		}
	}

	fmt.Println(h * d)
}

func parseFile() []move {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var values []move
	for _, v := range raw {
		parsed := strings.Split(v, " ")
		if err != nil {
			log.Fatalln(err)
		}

		steps, err := strconv.Atoi(parsed[1])
		if err != nil {
			log.Fatalln(err)
		}
		values = append(values, move{parsed[0], steps})
	}

	return values
}
