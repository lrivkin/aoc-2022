package main

import (
	"fmt"

	"github.com/lrivkin/aoc-2022/utils"
	"github.com/noahschumacher/go-collections/collections"
)

func parseInput(path string) [][]string {
	lines, _ := utils.ReadLines(path)
	sacks := make([][]string, len(lines))
	for i, l := range lines {
		first := l[0 : len(l)/2]
		second := l[len(l)/2:]
		sacks[i] = []string{first, second}
	}
	// fmt.Println(sacks)
	return sacks
}

func part1(in [][]string) int {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	weights := map[string]int{}
	for i := range letters {
		weights[letters[i:i+1]] = i + 1
	}
	// fmt.Println(weights)
	priority := 0
	for _, sack := range in {
		first := collections.Set[string]{}
		for i := range sack[0] {
			first.Add(sack[0][i : i+1])
		}
		second := collections.Set[string]{}
		for i := range sack[1] {
			second.Add(sack[1][i : i+1])
		}
		// fmt.Println(first)
		// fmt.Println(second)
		common := first.Intersection(second).ToSlice()[0]
		// fmt.Println(common)
		priority += weights[common]
		// fmt.Println(common[0] - 'A' + 1)
	}
	fmt.Printf("part1 = %d\n", priority)
	return priority
}

func part2(path string) int {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	weights := map[string]int{}
	for i := range letters {
		weights[letters[i:i+1]] = i + 1
	}
	priority := 0

	lines, _ := utils.ReadLines(path)
	common := collections.Set[string]{}
	for i, l := range lines {
		round := collections.Set[string]{}
		for j := range l {
			round.Add(l[j : j+1])
		}
		if i%3 == 0 {
			// initialize the set in common
			common = round
		}
		common = common.Intersection(round)
		if i%3 == 2 {
			// add to priority
			p := weights[common.ToSlice()[0]]
			// fmt.Printf("%d\t%d %v\n", i, p, common.ToSlice())
			priority += p
		}
	}
	fmt.Printf("part2 = %d\n", priority)
	return priority

}
func main() {
	test := parseInput("test.txt")
	part1(test)
	part1(parseInput("input.txt"))

	part2("test.txt")
	part2("input.txt")

}
