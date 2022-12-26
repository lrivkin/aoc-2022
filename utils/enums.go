package utils

import "fmt"

type Direction int

const (
	Down = iota
	Up
	Right
	Left
)

var (
	directionsMap = map[string]Direction{
		"D": Down,
		"U": Up,
		"R": Right,
		"L": Left,
	}
)

func (d Direction) String() string {
	switch d {
	case Down:
		return "D"
	case Up:
		return "U"
	case Left:
		return "L"
	case Right:
		return "R"
	default:
		return fmt.Sprintf("%d", int(d))
	}
}

func NewDirection(d string) (Direction, error) {
	direction, ok := directionsMap[d]
	if !ok {
		return -1, fmt.Errorf("invalid choice")
	}
	return direction, nil
}
