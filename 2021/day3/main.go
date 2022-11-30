package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	//problemOne()
	problemTwo()
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

func problemTwo() {
	origValues, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	binLen := len(origValues[0])

	var values []string
	values = append(values, origValues...)
	for i := 0; i < binLen; i++ {
		var mostCommon int
		var ones, zeros []string
		for _, v := range values {
			if string(v[i]) == "1" {
				ones = append(ones, v)
				mostCommon++
			} else {
				zeros = append(zeros, v)
				mostCommon--
			}
		}

		values = nil
		if mostCommon >= 0 {
			values = append(values, ones...)
		} else {
			values = append(values, zeros...)
		}

		if len(values) == 1 {
			break
		}
	}

	o2, err := strconv.ParseInt(values[0], 2, 64)
	if err != nil {
		log.Fatalln(err)
	}

	values = nil
	values = append(values, origValues...)
	for i := 0; i < binLen; i++ {
		var mostCommon int
		var ones, zeros []string
		for _, v := range values {
			if string(v[i]) == "1" {
				ones = append(ones, v)
				mostCommon++
			} else {
				zeros = append(zeros, v)
				mostCommon--
			}
		}

		values = nil
		if mostCommon < 0 {
			values = append(values, ones...)
		} else {
			values = append(values, zeros...)
		}

		if len(values) == 1 {
			break
		}
	}

	co2, err := strconv.ParseInt(values[0], 2, 64)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(o2 * co2)
}
