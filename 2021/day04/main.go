package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

type board [5][5]int

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	numbers, boards := parseFile()

	found := make([]board, len(boards))
	for _, num := range numbers {
		for i, board := range boards {
			var notFoundSum int
			var rowFound, columnFound [5]int

			for r, row := range board {
				for c, cv := range row {
					if found[i][r][c] == 1 || cv == num {
						found[i][r][c] = 1
						rowFound[r]++
						columnFound[c]++
					} else {
						notFoundSum += cv
					}
				}
			}

			for _, v := range append(rowFound[:], columnFound[:]...) {
				if v == 5 {
					fmt.Println(num * notFoundSum)
					return
				}
			}

		}
	}
}

func problemTwo() {
	numbers, boards := parseFile()

	found := make([]board, len(boards))
	winner := make([]bool, len(boards))
	var winners int
	for _, num := range numbers {
		for i, board := range boards {
			if winner[i] {
				continue
			}

			var notFoundSum int
			var rowFound, columnFound [5]int
			for r, row := range board {
				for c, cv := range row {
					if found[i][r][c] == 1 || cv == num {
						found[i][r][c] = 1
						rowFound[r]++
						columnFound[c]++
					} else {
						notFoundSum += cv
					}
				}
			}

			for _, v := range append(rowFound[:], columnFound[:]...) {
				if v == 5 {
					winner[i] = true
					winners++
					break
				}
			}

			if winners == len(boards) {
				fmt.Println(notFoundSum * num)
				return
			}
		}
	}
}

func parseFile() ([]int, []board) {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	numbersRaw := strings.Split(raw[0], ",")

	var numbers []int
	for _, numberRaw := range numbersRaw {
		number, err := strconv.Atoi(numberRaw)
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, number)
	}

	var boards []board
	for i := 2; i < len(raw); i += 6 {
		var b board
		for r := 0; r < 5; r++ {
			numbers := strings.Fields(raw[i+r])
			for c := 0; c < 5; c++ {
				b[r][c], err = strconv.Atoi(numbers[c])
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
		boards = append(boards, b)
	}

	return numbers, boards
}
