package main

import (
	"fmt"
	"os"
	"strings"
)

func part1(p1, p2 string, m map[string]int) int {
	roundPoints := m[p2]

	// (0 if you lost, 3 if the round was a draw, and 6 if you won)
	game := m[p2] - m[p1]
	switch game {
	case 0:
		// draw
		roundPoints += 3
	case 1, -2:
		// scissors > paper or paper > rock,
		roundPoints += 6
	}
	return roundPoints
}

func main() {
	fmt.Println("day2")
	file, _ := os.ReadFile("input.txt")
	rounds := strings.Split(string(file), "\n")
	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors
	wins := map[string]string{
		"A": "B",
		"B": "C",
		"C": "A",
	}
	loss := map[string]string{
		"A": "C",
		"B": "A",
		"C": "B",
	}
	part1Points := 0
	part2Points := 0
	for _, r := range rounds {

		// (1 for Rock, 2 for Paper, and 3 for Scissors)
		s := strings.Fields(r)
		p1 := s[0]
		p2 := s[1]
		part1Points += part1(p1, p2, m)

		p2round := 0
		myPlay := ""
		//  X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
		if p2 == "X" {
			// you lose
			myPlay = loss[p1]
		} else if p2 == "Y" {
			// draw
			p2round += 3
			myPlay = p1
		} else if p2 == "Z" {
			// win
			p2round += 6
			myPlay = wins[p1]
		}
		p2round += m[myPlay]
		// fmt.Printf("opponent= %s, outcome = %s, I choose %s, points= %d\n", p1, p2, myPlay, p2round)
		part2Points += p2round
	}
	fmt.Printf("part1: %d\n", part1Points)
	fmt.Printf("part2: %d\n", part2Points)

}
