package main

import (
	"reflect"
	"testing"
)

func TestGeneratePriorityMap(t *testing.T) {
	priorities := generatePriorityMap()

	testCases := []struct {
		letter   rune
		priority int
	}{
		{'a', 1},
		{'c', 3},
		{'A', 27},
		{'C', 29},
		{'z', 26},
		{'Z', 52},
		{'p', 16},
		{'L', 38},
		{'P', 42},
		{'v', 22},
		{'t', 20},
		{'s', 19},
	}

	for _, test := range testCases {
		if priorities[test.letter] != test.priority {
			t.Fatalf("(%s) expected: %d got: %d", string(test.letter), test.priority, priorities[test.letter])
		}
	}
}

func TestCommonItems(t *testing.T) {
	testCases := []struct {
		r      rucksack
		common []rune
	}{
		{rucksack{"AAA", "ABC"}, []rune{'A'}},
		{rucksack{"ABC", "AAA"}, []rune{'A'}},
		{rucksack{"abc", "aaa"}, []rune{'a'}},
		{rucksack{"aaa", "abc"}, []rune{'a'}},
		{rucksack{"AaAB", "AaAB"}, []rune{'A', 'a', 'B'}},
		{rucksack{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}, []rune{'p'}},
		{rucksack{"PmmdzqPrV", "vPwwTWBwg"}, []rune{'P'}},
		{rucksack{"zmmdzqzrV", "vPwwTWBwg"}, []rune{}},
	}

	for _, test := range testCases {
		common := commonItems(test.r)
		if len(common) != len(test.common) {
			t.Fatalf("mismatch %v, %v", common, test.common)
		}
	}

}

var example = []rucksack{
	{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
	{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
	{"PmmdzqPrV", "vPwwTWBwg"},
	{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
	{"ttgJtRGJ", "QctTZtZT"},
	{"CrZsJsPPZsGz", "wwsLwLmpwMDw"},
}

func TestTotalPriority(t *testing.T) {
	priority := totalPriority(example)
	if priority != 157 {
		t.Fatal("example mismatch")
	}
}

func TestParse(t *testing.T) {
	testParse := parseFile("./input_test.txt")
	if !reflect.DeepEqual(testParse, example) {
		t.Fatalf("parse mismatch:\n%v\n\n%v", example, testParse)
	}
}
