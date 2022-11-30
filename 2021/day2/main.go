package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	readFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var values []move
	for fileScanner.Scan() {

		parsed := strings.Split(fileScanner.Text(), " ")
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
