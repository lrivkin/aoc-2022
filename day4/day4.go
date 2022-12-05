package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/lrivkin/aoc-2022/utils"
)

func parseInput(path string) [][]int {
	re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
	file, _ := utils.ReadLines(path)
	nums := make([][]int, len(file))
	for i := range file {
		matches := re.FindStringSubmatch(file[i])
		ids := make([]int, 4)
		for j, m := range matches[1:] {
			val, _ := strconv.Atoi(m)
			ids[j] = val
		}
		nums[i] = ids
	}
	// fmt.Println(nums)
	return nums
}
func part1(ranges [][]int) int {
	contained := 0
	for _, i := range ranges {
		if (i[0] <= i[2] && i[1] >= i[3]) || (i[0] >= i[2] && i[1] <= i[3]) {
			contained += 1
			// fmt.Printf("ranges %v fully contained\n", i)
		}
	}
	fmt.Printf("part1 = %v\n", contained)
	return contained
}

func part2(ranges [][]int) int {
	lap := 0
	for _, i := range ranges {
		if (i[0] <= i[2] && i[1] >= i[2]) || (i[2] <= i[0] && i[3] >= i[0]) {
			lap += 1
			// fmt.Printf("ranges %v partially contained\n", i)
		}
	}
	fmt.Printf("part2 = %v\n", lap)
	return lap
}

func main() {
	part1(parseInput("test.txt"))
	part1(parseInput("input.txt"))
	part2(parseInput("test.txt"))
	part2(parseInput("input.txt"))
}
