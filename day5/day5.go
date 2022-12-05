package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(moves [][]int, blocks [][]string) string {
	for _, m := range moves {
		count := m[0]
		stackSrc := m[1]
		stackDest := m[2]

		for i := 0; i < count; i += 1 {
			movingBlock := len(blocks[stackSrc]) - 1 - i
			blocks[stackDest] = append(blocks[stackDest], blocks[stackSrc][movingBlock])
		}
		blocks[stackSrc] = blocks[stackSrc][0 : len(blocks[stackSrc])-count]
		// fmt.Printf("%v\n", blocks)
	}
	fmt.Printf("%v\n", blocks)
	topBlocks := ""
	for i := 1; i < len(blocks); i += 1 {
		topBlocks = fmt.Sprintf("%s%s", topBlocks, blocks[i][len(blocks[i])-1])
	}
	return topBlocks
}
func part2() int {
	return 0
}

func parseMoves(m string) [][]int {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	moves := re.FindAllStringSubmatch(m, -1)
	mp := make([][]int, len(moves))
	for i := range moves {
		mp[i] = make([]int, 3)
		for j, val := range moves[i][1:] {
			mp[i][j], _ = strconv.Atoi(val)
		}
	}
	return mp
}
func parseBlocks(in string) [][]string {
	// strings.Split(in, "\n")
	split := strings.Split(in, "\n")
	size := len(split[0]) / 4
	var blocks [][]string
	for i := 0; i <= size; i += 1 {
		blocks = append(blocks, []string{})
	}
	re := regexp.MustCompile(`\[(\w)\]`)
	// read in reverse order so we can use a "stack"
	for i := len(split) - 2; i >= 0; i -= 1 {
		line := split[i]
		for j := 0; j < size; j += 1 {
			block := line[4*j : 4*(j+1)]
			if re.MatchString(block) {
				blocks[j+1] = append(blocks[j+1], re.FindStringSubmatch(block)[1])
			}
		}
	}
	return blocks
}
func main() {
	file, _ := os.ReadFile("input.txt")
	p := strings.Split(string(file), "\n\n")
	moves := parseMoves(p[1])
	blocks := parseBlocks(p[0])
	print(part1(moves, blocks))

}
