package main

import (
	"fmt"

	"github.com/lrivkin/aoc-2022/utils"
	"github.com/noahschumacher/go-collections/collections"
)

func part1(in [][]int) int {
	cols := len(in[0])
	vis := collections.Set[int]{}

	for i := 0; i < len(in); i++ {
		row := in[i]
		// visible from left
		max := -1
		// fmt.Printf("checking row %d from the left\n", i)
		for j := 0; j < len(row); j++ {
			n := row[j]
			if n > max {
				// fmt.Printf("i=%d j=%d n=%d\n", i, j, n)
				vis.Add(i*cols + j)
				max = n
			}
		}
		// visible from right
		// fmt.Printf("checking row %d from the right\n", i)
		max = -1
		for j := len(row) - 1; j >= 0; j-- {
			n := row[j]
			if n > max {
				// fmt.Printf("i=%d j=%d n=%d\n", i, j, n)
				vis.Add(i*cols + j)
				max = n
			}
		}
	}
	for j := 0; j < cols; j++ {
		top := -1
		// fmt.Printf("checking col %d from the top \n", j)
		for i := 0; i < len(in); i++ {
			if in[i][j] > top {
				top = in[i][j]
				// fmt.Printf("i=%d j=%d n=%d\n", i, j, in[i][j])
				vis.Add(i*cols + j)
			}
		}

		bottom := -1
		// fmt.Printf("checking col %d from the bottom \n", j)
		for i := cols - 1; i >= 0; i-- {
			if in[i][j] > bottom {
				bottom = in[i][j]
				// fmt.Printf("i=%d j=%d n=%d\n", i, j, in[i][j])
				vis.Add(i*cols + j)
			}
		}
	}
	// fmt.Printf("visible= %d, all positions: %v\n", len(vis), vis.ToSlice())
	fmt.Printf("visible= %d\n", len(vis))
	return len(vis)
}

func part2(in [][]int) int {
	bestView := 0
	for i := 1; i < len(in)-1; i++ {
		for j := 1; j < len(in[0])-1; j++ {
			// look up/down/left/right
			up := i + 1
			down := i - 1
			left := j - 1
			right := j + 1
			val := in[i][j]
			for down > 0 {
				if in[down][j] >= val {
					break
				}
				down--
			}
			for up < len(in)-1 {
				if in[up][j] >= val {
					break
				}
				up++
			}
			for left > 0 {
				if in[i][left] >= val {
					break
				}
				left--
			}
			for right < len(in[0])-1 {
				if in[i][right] >= val {
					break
				}
				right++
			}
			score := (up - i) * (i - down) * (j - left) * (right - j)
			// fmt.Printf("(%d,%d) val=%d, {down:%d, left:%d, up:%d, right:%d} score = %d\n",
			// 	i, j, in[i][j], (up - i), (j - left), (i - down), (right - j), score)
			if score > bestView {
				bestView = score
			}
		}
	}
	fmt.Printf("part2 = %d\n", bestView)
	return bestView
}
func main() {
	nums, _ := utils.ParseIntGrid("test.txt")
	part1(nums)
	part2(nums)

	nums, _ = utils.ParseIntGrid("input.txt")
	part1(nums)
	part2(nums)
}
