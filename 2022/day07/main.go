package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/chrishoffman/advent-of-code"
)

type file struct {
	name string
	size int
}

type directory struct {
	directories map[string]directory
	files       []file
	size        int
}

func main() {
	problemOne()
	problemTwo()
}

func problemOne() {
	var total int

	var f func(directory)
	f = func(dir directory) {
		for _, d := range dir.directories {
			f(d)
		}
		if dir.size <= 100000 {
			total += dir.size
		}
	}

	dir := parseFile()
	f(dir)

	fmt.Println(total)
}

func problemTwo() {
	var candidateDir []int

	dir := parseFile()
	spaceAvalable := 70000000 - dir.size
	spaceRequired := 30000000 - spaceAvalable

	var f func(directory)
	f = func(dir directory) {
		for _, d := range dir.directories {
			f(d)
		}
		if dir.size >= spaceRequired {
			candidateDir = append(candidateDir, dir.size)
		}
	}

	f(dir)

	sort.Ints(candidateDir)
	fmt.Println(candidateDir[0])
}

func parseFile() directory {
	raw, err := advent.ParseFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	dir, _ := parseDir(raw, 2)
	return dir
}

func parseDir(raw []string, startIndex int) (directory, int) {
	dir := directory{
		files:       make([]file, 0),
		directories: make(map[string]directory),
	}

	for i := startIndex; i < len(raw); i++ {
		v := strings.Fields(raw[i])
		switch v[0] {
		case "$":
			if v[1] == "cd" {
				if v[2] == ".." {
					return dir, i
				}
				i += 1
				nested, newIndex := parseDir(raw, i)
				dir.directories[v[2]] = nested
				dir.size += nested.size
				i = newIndex
			}
		case "dir":
			continue
		default:
			size, _ := strconv.Atoi(v[0])
			dir.files = append(dir.files, file{name: v[1], size: size})
			dir.size += size
		}
	}
	return dir, len(raw)
}
