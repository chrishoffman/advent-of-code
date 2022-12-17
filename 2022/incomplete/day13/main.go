package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	input := parseFile()

NEXT:
	for _, i := range input {
		left := i.left.unwrap()
		right := i.right.unwrap()

		
	}

}

func problemTwo() {
}

type signal string

func (p signal) unwrap() string {
	var sb strings.Builder
	for i := 1; i < len(p)-1; i++ {
		switch p[i] {
		case '[':
			nested := fmt.Sprintf("[%s]", signal(p[i:]).unwrap())
			sb.WriteString(nested)
			i += len(nested) - 1
		case ']':
			return sb.String()
		default:
			sb.WriteByte(p[i])
		}
	}
	return sb.String()
}

type input struct{ left, right signal }

func (i input) next() { //([]int, []int, input) {
	// left
	left := i.left.unwrap()
	right := i.right.unwrap()
	for {

	}

	fmt.Println(left, right)
}

func parseFile() []input {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	var d []input
	for i := 0; i < len(raw); i += 3 {
		d = append(d, input{signal(raw[i]), signal(raw[i+1])})
	}

	return d
}
