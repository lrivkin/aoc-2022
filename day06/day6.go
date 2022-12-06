package main

import (
	"fmt"

	"github.com/lrivkin/aoc-2022/utils"
)

func part1(in string) int {
	var seen map[string]struct{}
	for i := 0; i < len(in); i += 1 {
		seen = map[string]struct{}{}
		for j := i; j < len(in) && j < i+14; j += 1 {
			_, ok := seen[in[j:j+1]]
			if ok {
				break
			}
			if j == i+13 {
				fmt.Printf("found it! %d, %s\n", j+1, in[i:j+1])
				return j + 1
			}
			seen[in[j:j+1]] = struct{}{}
		}

	}
	// fmt.Println(chars)
	return 0
}

func main() {
	tests, _ := utils.ReadBlock("test.txt")
	for _, t := range tests {
		part1(t[0])
	}
	in, _ := utils.ReadLines("input.txt")
	part1(in[0])
}
