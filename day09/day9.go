package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lrivkin/aoc-2022/utils"
)

type Move struct {
	Direction utils.Direction
	Num       int
}

type Position struct {
	x, y int
}

func NewPosition() *Position {
	return &Position{0, 0}
}

func (p *Position) move(d utils.Direction) {
	switch d {
	case utils.Down:
		p.y--
	case utils.Up:
		p.y++
	case utils.Left:
		p.x--
	case utils.Right:
		p.x++
	}
}

type Snake struct {
	head, tail *Position
}

func (s Snake) String() string {
	if s.head != nil && s.tail != nil {
		return fmt.Sprintf("%v-%v", *s.head, *s.tail)
	}
	return fmt.Sprintf("%v - %v", s.head, s.tail)
}

func NewSnake() *Snake {
	return &Snake{
		head: NewPosition(),
		tail: NewPosition(),
	}
}

// Return the positions the tail went on this turn
func (s *Snake) move(m Move) []Position {
	positions := []Position{*s.tail}
	for i := 0; i < m.Num; i++ {
		s.head.move(m.Direction)
		if utils.AbsVal(s.head.x-s.tail.x) > 1 {
			s.tail.y = s.head.y
			s.tail.move(m.Direction)
		} else if utils.AbsVal(s.head.y-s.tail.y) > 1 {
			s.tail.x = s.head.x
			s.tail.move(m.Direction)
		}
		positions = append(positions, *s.tail)
		// fmt.Printf("%v\n", s)
	}
	return positions
}

type Rope struct {
	knots []*Position
}

func NewRope() *Rope {
	positions := make([]*Position, 10)
	for i := 0; i < 10; i++ {
		positions[i] = NewPosition()
	}
	return &Rope{positions}
}

func (r Rope) String() string {
	ret := fmt.Sprintf("%v", *r.knots[0])
	for i := 1; i < len(r.knots); i++ {
		ret = fmt.Sprintf("%s-%v", ret, *r.knots[i])
	}
	return ret
}

func (r *Rope) move(m Move) []Position {
	// fmt.Println(m)
	tail := r.knots[9]
	tailPos := *tail
	positions := []Position{tailPos}
	for i := 0; i < m.Num; i++ {
		ahead := r.knots[0]
		ahead.move(m.Direction)
		for j := 1; j < len(r.knots); j++ {
			cur := r.knots[j]
			ahead := r.knots[j-1]
			// right
			if ahead.x-cur.x > 1 {
				if ahead.y > cur.y {
					cur.move(utils.Up)
				} else if ahead.y < cur.y {
					cur.move(utils.Down)
				}
				cur.move(utils.Right)
			}
			// left
			if ahead.x-cur.x < -1 {
				if ahead.y > cur.y {
					cur.move(utils.Up)
				} else if ahead.y < cur.y {
					cur.move(utils.Down)
				}
				cur.move(utils.Left)
			}
			// up
			if ahead.y-cur.y > 1 {
				if ahead.x > cur.x {
					cur.move(utils.Right)
				} else if ahead.x < cur.x {
					cur.move(utils.Left)
				}
				cur.move(utils.Up)
			}
			// down
			if ahead.y-cur.y < -1 {
				if ahead.x > cur.x {
					cur.move(utils.Right)
				} else if ahead.x < cur.x {
					cur.move(utils.Left)
				}
				cur.move(utils.Down)
			}
		}
		if *tail != tailPos {
			positions = append(positions, *tail)
			tailPos.x = tail.x
			tailPos.y = tail.y
		}
		// fmt.Printf("%v\n", r)
	}
	// fmt.Printf("snake: %v\n", r)
	// fmt.Printf("tails: %v\n", positions)
	return positions
}

type Grid map[Position]struct{}

func NewGrid() Grid {
	return make(Grid, 0)
}

func (g Grid) add(p Position) {
	g[p] = struct{}{}
}

func (grid Grid) String() string {
	spots := make([]string, len(grid))
	i := 0
	for g := range grid {
		spots[i] = fmt.Sprintf("(%d,%d)", g.x, g.y)
		i++
	}
	return fmt.Sprintf("Grid: Len=%d, values=[%s]", len(grid), strings.Join(spots, " "))
}

func parse(path string) ([]Move, error) {
	lines, err := utils.ReadLines(path)
	if err != nil {
		return nil, err
	}
	moves := make([]Move, len(lines))
	for i, l := range lines {
		split := strings.Split(l, " ")
		n, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		direction, err := utils.NewDirection(split[0])
		if err != nil {
			return nil, err
		}
		moves[i] = Move{direction, n}
	}
	return moves, nil
}

func part1(moves []Move) int {
	snake := NewSnake()
	grid := NewGrid()
	// fmt.Printf("%v\n", snake)
	for _, move := range moves {
		tails := snake.move(move)
		for _, t := range tails {
			grid.add(t)
		}
	}
	fmt.Printf("Part 1: Total tail spots = %d\n", len(grid))
	return len(grid)
}

func part2(moves []Move) int {
	rope := NewRope()
	grid := NewGrid()
	// fmt.Printf("%v\n", rope)
	for _, move := range moves {
		tails := rope.move(move)
		for _, t := range tails {
			grid.add(t)
		}
	}
	// fmt.Printf("all tail spots = %d (", len(grid))
	// fmt.Println(grid)
	fmt.Printf("Part 2: New tail spots = %d\n", len(grid))
	return len(grid)
}

func main() {
	test, err := parse("test.txt")
	if err != nil {
		panic(err)
	}
	part1(test)
	part2(test)
	test2, _ := parse("test2.txt")
	part2(test2)
	fmt.Printf("\nMy input\n")
	input, _ := parse("input.txt")
	part1(input)
	part2(input)
}
