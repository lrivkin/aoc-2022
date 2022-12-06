package main

import (
	"fmt"

	"github.com/lrivkin/aoc-2022/utils"
)

func findIdx(in string, numChars int) int {
	seen := make(map[string]int, numChars)
	minIdx := 0
	for i := 0; i < len(in); i += 1 {
		char := in[i : i+1]
		oldIdx, ok := seen[char]
		if ok {
			for j := minIdx; j <= oldIdx; j += 1 {
				delete(seen, in[j:j+1])
			}
			minIdx = oldIdx + 1
		}
		seen[char] = i
		if i-minIdx > numChars {
			delete(seen, in[minIdx:minIdx+1])
			minIdx += 1
		}
		// fmt.Printf("%v\n", seen)
		if len(seen) == numChars {
			fmt.Printf("%d unique characters\t%d %s\n", numChars, i+1, in[i-numChars:i+1])
			return i + 1
		}
	}
	return 0
}

func main() {
	fmt.Println("Tests")
	tests, _ := utils.ReadLines("test.txt")
	for _, t := range tests {
		findIdx(t, 4)
		findIdx(t, 14)

	}
	fmt.Printf("\nMy input\n")
	in, _ := utils.ReadLines("input.txt")
	findIdx(in[0], 4)
	findIdx(in[0], 14)

}
