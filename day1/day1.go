package main

import (
	"fmt"
	"sort"

	"github.com/lrivkin/aoc-2022/utils"
)

func parseIn(file string) [][]int {
	inputs, _ := utils.ReadBlock(file)
	nums := make([][]int, len(inputs))
	for i, x := range inputs {
		n, _ := utils.StringSliceToIntSlice(x)
		nums[i] = n
	}
	return nums
}

func part1(in [][]int) int {
	totalCals := make([]int, len(in))
	for i, x := range in {
		s := utils.Sum(x)
		// fmt.Printf("total=%d, %v\n", s, x)
		totalCals[i] = s
	}
	return utils.Max(totalCals)
}

func part2(in [][]int) int {
	totalCals := make([]int, len(in))
	for i, x := range in {
		s := utils.Sum(x)
		// fmt.Printf("total=%d, %v\n", s, x)
		totalCals[i] = s
	}
	sort.Ints(totalCals)
	return utils.Sum(totalCals[len(totalCals)-3:])
}
func main() {
	test := parseIn("test.txt")
	fmt.Printf("test: part1=%d\n", part1(test))
	fmt.Printf("test: part2=%d\n", part2(test))

	input := parseIn("day1-input.txt")
	fmt.Printf("part1=%d\n", part1(input))
	fmt.Printf("part2=%d\n", part2(input))

}
