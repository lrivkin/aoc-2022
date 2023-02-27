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

var LCM uint64 = 1

type Monkey struct {
	items      []uint64
	op         *ast.BinaryExpr
	testVal    uint64
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
			m.items = make([]uint64, len(items))
			for i, x := range items {
				m.items[i] = uint64(x)
			}
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
			val, _ := strconv.ParseUint(strings.TrimPrefix(strings.TrimSpace(line), "Test: divisible by "), 10, 64)
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

func (m *Monkey) applyOp(item uint64) uint64 {
	// fmt.Printf("X: %v\tOP: %v\tY: %v\n", m.op.X, m.op.Op, m.op.Y)
	var y uint64
	switch m.op.Y.(type) {
	case *ast.Ident:
		y = item
	case *ast.BasicLit:
		y, _ = strconv.ParseUint(m.op.Y.(*ast.BasicLit).Value, 10, 64)
	}
	switch m.op.Op {
	case token.ADD:
		// fmt.Printf("%d + %d\n", item, y)
		return ((item % LCM) + (y % LCM)) % LCM
	case token.MUL:
		// fmt.Printf("%d * %d\n", item, y)
		return ((item % LCM) * (y % LCM)) % LCM
	default:
		return 0
	}
}

func (m *Monkey) tossItem(item uint64) (uint64, int) {
	// apply function
	// fmt.Printf("start: item= %d\n", item)
	item = m.applyOp(item)
	// fmt.Printf("function applied: item= %d\n", item)

	// item = item / 3
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
	for _, m := range monkeys {
		LCM *= m.testVal
	}
	// fmt.Println("start")
	// printUtil(monkeys)
	for r := 0; r < 10000; r++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				nevVal, tossTo := m.tossItem(item)
				monkeys[tossTo].items = append(monkeys[tossTo].items, nevVal)
			}
			m.numItems += len(m.items)
			m.items = m.items[0:0]
		}
		// fmt.Printf("After round %d\n", r+1)
		// printUtil(monkeys)
		if (r+1)%1000 == 0 || r == 19 || r == 0 {
			fmt.Printf("\n== After round %d ==\n", r+1)
			for i, m := range monkeys {
				fmt.Printf("Monkey %d inspected items %d times.\n", i, m.numItems)
			}
		}
	}
	tosses := make([]int, len(monkeys))
	for i := range monkeys {
		tosses[i] = monkeys[i].numItems
	}
	utils.SortSlice(tosses, true)
	return tosses[0] * tosses[1]
}

func main() {
	test := parseMonkeys("input.txt")
	fmt.Println(part1(test))
}
