package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	//problemTwo()
}

func problemOne() {
	values, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	mostCommon := make([]int, len(values[0]))
	for _, v := range values {
		for i, h := range v {
			if string(h) == "1" {
				mostCommon[i]++
			} else {
				mostCommon[i]--
			}
		}
	}

	var gammaRaw, epsilonRaw string
	for _, c := range mostCommon {
		if c > 0 {
			gammaRaw += "1"
			epsilonRaw += "0"
		} else {
			gammaRaw += "0"
			epsilonRaw += "1"
		}
	}

	gamma, err := strconv.ParseInt(gammaRaw, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}
	epsilon, err := strconv.ParseInt(epsilonRaw, 2, 64)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(gamma * epsilon)
}

// func problemTwo() {
// 	values, err := advent.ParseFile("./input.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	newValues :=
// 	mostCommon := make([]int, len(values[0]))
// 	for i, h := range v {

// 		for _, v := range values {
// 			for i, h := range v {
// 				if string(h) == "1" {
// 					mostCommon[i]++
// 				} else {
// 					mostCommon[i]--
// 				}
// 			}
// 	}

// }
