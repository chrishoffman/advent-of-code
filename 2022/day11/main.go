package main

import (
	"log"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
}

func problemTwo() {
}

type monkey struct {
	items      []int
	test       func(int) bool
	operatiion func(int) int
	action     map[bool]int
}

func parseFile() []monkey {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

}
