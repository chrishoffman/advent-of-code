package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	forest := parseFile()

	var visible int
	for r := 0; r < len(forest); r++ {
		for c := 0; c < len(forest[r]); c++ {
			height := forest[r][c]
			var sidesHidden int
			for u := r - 1; u >= 0; u-- {
				if forest[u][c] >= height {
					sidesHidden++
					break
				}
			}
			for d := r + 1; d < len(forest); d++ {
				if forest[d][c] >= height {
					sidesHidden++
					break
				}
			}
			for l := c - 1; l >= 0; l-- {
				if forest[r][l] >= height {
					sidesHidden++
					break
				}
			}
			for i := c + 1; i < len(forest[r]); i++ {
				if forest[r][i] >= height {
					sidesHidden++
					break
				}
			}
			if sidesHidden < 4 {
				visible++
			}
		}
	}

	fmt.Println(visible)
}

func problemTwo() {
	forest := parseFile()

	var maxScenicScore int
	for r := 0; r < len(forest); r++ {
		for c := 0; c < len(forest[r]); c++ {
			height := forest[r][c]

			var up, down, left, right int
			for u := r - 1; u >= 0; u-- {
				up++
				if forest[u][c] >= height {
					break
				}
			}
			for d := r + 1; d < len(forest); d++ {
				down++
				if forest[d][c] >= height {
					break
				}
			}
			for l := c - 1; l >= 0; l-- {
				left++
				if forest[r][l] >= height {
					break
				}
			}
			for i := c + 1; i < len(forest[r]); i++ {
				right++
				if forest[r][i] >= height {
					break
				}
			}

			scenicScore := up * down * left * right
			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	fmt.Println(maxScenicScore)
}

func parseFile() [][]int {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	forest := make([][]int, len(raw))
	for r := 0; r < len(raw); r++ {
		forest[r] = make([]int, len(raw[r]))
		for c := 0; c < len(raw[r]); c++ {
			forest[r][c], _ = strconv.Atoi(string(raw[r][c]))
		}
	}

	return forest
}
