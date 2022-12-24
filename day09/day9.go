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

func part2() int {
	return 0
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
		return fmt.Sprintf("%v - %v", *s.head, *s.tail)
	}
	return fmt.Sprintf("%v - %v", s.head, s.tail)
}

func NewSnake() *Snake {
	return &Snake{
		head: NewPosition(),
		tail: NewPosition(),
	}
}

type Grid map[Position]struct{}

func NewGrid() Grid {
	return make(Grid, 0)
}

func (g Grid) add(p Position) {
	g[p] = struct{}{}
}

// Return the positions the tail went on this turn
func (s *Snake) move(m Move) []Position {
	positions := []Position{*s.tail}
	fmt.Printf("%v ", *s.tail)
	for i := 0; i < m.Num; i++ {
		s.head.move(m.Direction)
	}
	// if the tail moves at all, do the direction one first
	switch m.Direction {
	case utils.Up, utils.Down:
		if utils.AbsVal(s.head.y-s.tail.y) > 1 {
			s.tail.x = s.head.x
			s.tail.move(m.Direction)
			fmt.Printf("%v ", *s.tail)
		}
	case utils.Left, utils.Right:
		if utils.AbsVal(s.head.x-s.tail.x) > 1 {
			s.tail.y = s.head.y
			s.tail.move(m.Direction)
			fmt.Printf("%v ", *s.tail)
		}
	}
	positions = append(positions, *s.tail)
	// finish in the straight line
	for i := 0; i < m.Num-2; i++ {
		s.tail.move(m.Direction)
		fmt.Printf("%v ", *s.tail)
		positions = append(positions, *s.tail)
	}
	fmt.Println()
	return positions
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
	fmt.Printf("%v\n", *snake)
	for _, move := range moves {
		tails := snake.move(move)
		for _, t := range tails {
			grid.add(t)
		}
		fmt.Printf("Move: %v, Snake: %v\n", move, *snake)
		// for _, g := range tails {
		// 	fmt.Printf("%v ", g)
		// }
		// fmt.Println()
	}
	return 0
}

func main() {
	test, err := parse("test.txt")
	if err != nil {
		panic(err)
	}
	part1(test)
	// fmt.Printf("%v\n", test)

}
