package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lrivkin/aoc-2022/utils"
)

func parseInput(path string) []int {
	parsed := []int{0}
	file, _ := utils.ReadLines(path)
	for _, line := range file {
		if line == "noop" {
			parsed = append(parsed, 0)
		} else {
			num, _ := strconv.Atoi(strings.Split(line, " ")[1])
			parsed = append(parsed, 0, num)
		}
	}
	return parsed
}

func part1(in []int) int {
	ans := 0
	i := 1
	x := 1
	for i <= 220 {
		if (i-20)%40 == 0 {
			fmt.Printf("%d * %d = %d\n", i, x, i*x)
			ans += i * x
		}
		// fmt.Printf("%d: x=%d, x+=%d\n", i, x, in[i])
		x += in[i]
		i++
	}
	fmt.Printf("part1 = %d\n", ans)
	return ans
}
func part2(in []int) int {
	// ans := 0
	crt := []string{}
	i := 1
	dot := 0
	x := 1
	row := ""
	for i <= 240 {
		cursorDiff := (dot - x)
		if cursorDiff >= -1 && cursorDiff <= 1 {
			row += "#"
		} else {
			row += "."
		}
		x += in[i]
		i++
		dot++
		if dot == 40 {
			fmt.Println(row)
			crt = append(crt, strings.Clone(row))
			row = ""
			dot = 0
		}

	}
	return 0
}
func main() {
	test := parseInput("test.txt")
	data := parseInput("input.txt")

	fmt.Printf("%#v\n", test)
	part1(test)
	fmt.Println()
	part1(data)

	part2(test)
	fmt.Println()
	part2(data)
}
