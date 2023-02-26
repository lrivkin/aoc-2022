package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"

	"github.com/lrivkin/aoc-2022/utils"
)

type Monkey struct {
	items      []int
	op         *ast.BinaryExpr
	testVal    int
	nextMonkey []int
	numItems   int
}

func NewMonkey(s string) *Monkey {
	m := &Monkey{
		nextMonkey: make([]int, 2),
	}
	for i, line := range strings.Split(s, "\n") {
		if i == 1 {
			rawItems := strings.TrimPrefix(strings.TrimSpace(line), "Starting items: ")
			// fmt.Println(rawItems)
			items, _ := utils.StringListToIntSlice(rawItems)
			m.items = items
		}
		if i == 2 {
			rawOp := strings.TrimPrefix(strings.TrimSpace(line), "Operation: new = ")
			// fmt.Println(rawOp)
			exp, err := parser.ParseExpr(rawOp)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			m.op = exp.(*ast.BinaryExpr)
		}
		if i == 3 {
			val, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSpace(line), "Test: divisible by "))
			m.testVal = val
		}
		if i == 4 {
			next, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSpace(line), "If true: throw to monkey "))
			m.nextMonkey[0] = next
		}
		if i == 5 {
			next, _ := strconv.Atoi(strings.TrimPrefix(strings.TrimSpace(line), "If false: throw to monkey "))
			m.nextMonkey[1] = next
		}
	}
	// fmt.Printf("%#v\n", m)
	return m
}
func (m *Monkey) String() string {
	return fmt.Sprintf("%v", m.items)
}

func printUtil(in []*Monkey) {
	for i, m := range in {
		fmt.Printf("Monkey %d: %s\n", i, m)
	}
	fmt.Println()
}

func (m *Monkey) applyOp(item int) int {
	// fmt.Printf("X: %v\tOP: %v\tY: %v\n", m.op.X, m.op.Op, m.op.Y)
	var y int
	switch m.op.Y.(type) {
	case *ast.Ident:
		y = item
	case *ast.BasicLit:
		y, _ = strconv.Atoi(m.op.Y.(*ast.BasicLit).Value)
	}
	switch m.op.Op {
	case token.ADD:
		// fmt.Printf("%d + %d\n", item, y)
		return item + y
	case token.MUL:
		// fmt.Printf("%d * %d\n", item, y)
		return item * y
	default:
		return -1
	}
}

func (m *Monkey) tossItem(item int) (int, int) {
	// apply function
	// fmt.Printf("start: item= %d\n", item)
	item = m.applyOp(item)
	// fmt.Printf("function applied: item= %d\n", item)

	item = item / 3
	// fmt.Printf("divided by 3: item= %d\n", item)

	if item%m.testVal == 0 {
		// fmt.Printf("divisible by %d, tossing to %d\n", m.testVal, m.nextMonkey[0])
		return item, m.nextMonkey[0]
	} else {
		// fmt.Printf("NOT divisible by %d, tossing to %d\n", m.testVal, m.nextMonkey[1])
		return item, m.nextMonkey[1]
	}
}
func parseMonkeys(path string) []*Monkey {
	monkeys := []*Monkey{}
	contents, _ := os.ReadFile(path)
	chunks := strings.Split(string(contents), "\n\n")
	fmt.Printf("got %d monkeys\n", len(chunks))
	for _, x := range chunks {
		monkeys = append(monkeys, NewMonkey(x))
	}
	return monkeys
}

func part1(monkeys []*Monkey) int {
	fmt.Println("start")
	printUtil(monkeys)
	for r := 0; r < 20; r++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				nevVal, tossTo := m.tossItem(item)
				monkeys[tossTo].items = append(monkeys[tossTo].items, nevVal)
			}
			m.numItems += len(m.items)
			m.items = m.items[0:0]
		}
		fmt.Printf("After round %d\n", r+1)
		printUtil(monkeys)
	}
	for i, m := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, m.numItems)
	}
	return -1
}
func part2() int {
	return -1
}
func main() {
	test := parseMonkeys("input.txt")
	part1(test)
}
