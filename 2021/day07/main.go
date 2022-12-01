package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	points, _ := parseFile()

	minPoint := points[0]
	maxPoint := points[len(points)-1]

	prevDistance := totalDistance(points, minPoint)
	for i := minPoint + 1; i < maxPoint; i++ {
		dist := totalDistance(points, i)
		if prevDistance < dist {
			break
		}
		prevDistance = dist
	}
	fmt.Println(prevDistance)
}

func problemTwo() {
	points, _ := parseFile()

	minPoint := points[0]
	maxPoint := points[len(points)-1]

	prevDistance := increasingCostTotalDistance(points, minPoint)
	for i := minPoint + 1; i < maxPoint; i++ {
		dist := increasingCostTotalDistance(points, i)
		if prevDistance < dist {
			break
		}
		prevDistance = dist
	}
	fmt.Println(prevDistance)
}

func totalDistance(points []int, dest int) int {
	var sum float64
	for _, p := range points {
		sum += math.Abs(float64(p) - float64(dest))
	}
	return int(sum)
}

func increasingCostTotalDistance(points []int, dest int) int {
	var sum float64
	for _, p := range points {
		dist := math.Abs(float64(p) - float64(dest))
		sum += (dist * (dist + 1)) / 2 // https://math.stackexchange.com/questions/593318/factorial-but-with-addition
	}
	return int(sum)
}

// returns sorted list of points and the sum
func parseFile() ([]int, int) {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	numbersRaw := strings.Split(raw[0], ",")
	var numbers []int
	var sum int
	for _, numberRaw := range numbersRaw {
		number, err := strconv.Atoi(numberRaw)
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, number)
		sum += number
	}
	sort.Ints(numbers)
	return numbers, sum
}
