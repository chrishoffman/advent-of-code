package main

import (
	"fmt"
	"log"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	distinctSearch(4)
}

func problemTwo() {
	distinctSearch(14)
}

func distinctSearch(l int) {
	data := parseFile()
	for b := 0; b < len(data); b++ {
		unique := make(map[byte]bool)
		for i := 0; i < l; i++ {
			unique[data[b+i]] = true
		}
		if len(unique) == l {
			fmt.Println(b + l)
			break
		}
	}
}

func parseFile() string {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return raw[0]
}
