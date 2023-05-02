package main

import (
	"fmt"
	"math"

	"github.com/lrivkin/aoc-2022/utils"
)

const (
	S = rune('S')
	E = rune('E')
)

type coordinate struct {
	x, y int
}

func (c coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.x, c.y)
}

type graph struct {
	edges map[coordinate][]coordinate
}

type hill struct {
	start, end coordinate
	heightmap  [][]rune
	graph      *graph
}

func (h *hill) String() string {
	s := fmt.Sprintf("Start=%v\tEnd=%v",
		h.start, h.end)
	for _, line := range h.heightmap {
		s = fmt.Sprintf("%s\n%v", s, line)
	}
	return s
}

func toHill(path string) *hill {
	lines, _ := utils.ReadLines(path)
	hill := &hill{
		heightmap: make([][]rune, len(lines)),
	}
	for i, l := range lines {
		hill.heightmap[i] = []rune(l)
		for j, val := range hill.heightmap[i] {
			switch val {
			case S:
				hill.start = coordinate{i, j}
				hill.heightmap[i][j] = rune('a')
			case E:
				hill.end = coordinate{i, j}
				hill.heightmap[i][j] = rune('z')
			}
		}
	}

	fmt.Printf("%v\n", hill)
	return hill
}

func (h *hill) buildGraph() {
	g := &graph{
		edges: map[coordinate][]coordinate{},
	}
	for i, row := range h.heightmap {
		for j, v := range row {
			c := coordinate{i, j}

			reachable := []coordinate{}

			for _, n := range []coordinate{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}} {
				if c == h.end {
					break
				}
				if n.x < 0 || n.x >= len(h.heightmap) || n.y < 0 || n.y >= len(row) {
					continue
				}
				if h.heightmap[i][j] == S || h.heightmap[i][j] == E ||
					h.heightmap[n.x][n.y]-v <= 1 {

					reachable = append(reachable, coordinate{n.x, n.y})
				}
			}
			g.edges[c] = reachable
			// fmt.Printf("%v: %c -> %v\n", c, v, reachable)
		}
	}
	h.graph = g
}

func (h *hill) doAlgo() int {
	q := map[coordinate]struct{}{}
	var empty struct{}

	// initialize distances
	dist := map[coordinate]int{}
	for e := range h.graph.edges {
		if h.heightmap[e.x][e.y] == rune('a') {
			dist[e] = 0
		} else {
			dist[e] = math.MaxInt
		}
		q[e] = empty
	}

	for len(q) > 0 {
		// pop node w/ the smallest distance
		min := math.MaxInt
		var v coordinate
		for c := range q {
			if dist[c] <= min {
				min = dist[c]
				v = c
			}
		}

		delete(q, v)
		// fmt.Printf("Next node = %v: %c min = %d\n", v, h.heightmap[v.x][v.y], min)
		if v == h.end {
			fmt.Printf("Found E @ %v! Shortest path = %d\n", h.end, min)
			break
		}

		for _, u := range h.graph.edges[v] {

			if _, ok := q[u]; !ok {
				// fmt.Printf("\tskipped:      %c %v\n", h.heightmap[u.x][u.y], u)
				continue
			}

			alt := dist[v] + 1

			if alt < dist[u] {
				dist[u] = alt
				// fmt.Printf("\tnew distance: %c %v = %d\n", h.heightmap[u.x][u.y], u, alt)
			} else {
				// fmt.Printf("\tunchanged:    %c %v\n", h.heightmap[u.x][u.y], u)
			}
		}

	}
	return dist[h.end]
}

func part2() int {
	return 0
}

func main() {
	h := toHill("input.txt")
	fmt.Println()
	h.buildGraph()
	h.doAlgo()
}
