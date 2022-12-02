package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

type play int

const (
	rock     play = 1
	paper         = 2
	scissors      = 3
)

var playMapping = map[string]play{
	"A": rock,
	"B": paper,
	"C": scissors,
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

type round struct{ opp, you play }

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	rounds := parseFile()
	var total int
	for _, r := range rounds {
		total += scoreHand(r)
	}
	fmt.Println(total)
}

func problemTwo() {
	rounds := parseFile()
	var total int
	for _, r := range rounds {
		transRound := translate(r)
		total += scoreHand(transRound)
	}
	fmt.Println(total)
}

func scoreHand(round round) int {
	score := int(round.you)

	if round.opp == round.you {
		return score + 3
	}

	if (round.you == rock && round.opp == scissors) ||
		(round.you == paper && round.opp == rock) ||
		(round.you == scissors && round.opp == paper) {
		score += 6
	}
	return score
}

// rock=loss, paper=draw, scissors=win
func translate(rnd round) round {
	if rnd.you == paper {
		return round{rnd.opp, rnd.opp}
	}

	var you play
	switch rnd.opp {
	case rock:
		if rnd.you == rock {
			you = scissors
		} else {
			you = paper
		}
	case scissors:
		if rnd.you == rock {
			you = paper
		} else {
			you = rock
		}
	case paper:
		if rnd.you == rock {
			you = rock
		} else {
			you = scissors
		}
	}
	return round{rnd.opp, you}
}

func parseFile() []round {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var rounds []round
	for _, l := range raw {
		plays := strings.Fields(l)

		rounds = append(rounds, round{playMapping[plays[0]], playMapping[plays[1]]})
	}
	return rounds
}
