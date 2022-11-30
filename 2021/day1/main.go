package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	values := parseFile()

	increaseCount := 0
	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			increaseCount++
		}
	}
	fmt.Println(increaseCount)
}

func problemTwo() {
	values := parseFile()

	increaseCount := 0
	for i := 1; i < len(values)-2; i++ {
		sum := values[i] + values[i+1] + values[i+2]
		prevSum := values[i-1] + values[i] + values[i+1]
		if sum > prevSum {
			increaseCount++
		}
	}
	fmt.Println(increaseCount)
}

func parseFile() []int {
	readFile, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var values []int
	for fileScanner.Scan() {
		value, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		values = append(values, value)
	}

	return values
}
