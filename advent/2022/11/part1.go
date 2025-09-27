package advent

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

type Monkey struct {
	Id         int
	Items      []int
	PlusOrMult bool
	Same       bool
	By         int
	Test       int
	TestTrue   int
	TestFalse  int
	Inspected  int
}

func StrsToIns(strs []string) []int {
	conv := []int{}
	for _, s := range strs {
		c, _ := strconv.Atoi(s)
		conv = append(conv, c)
	}
	return conv
}

func Year20221101() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2022/11/ie.txt")
	lines := advent.GetLinesFromFile("advent/2022/11/input.txt")

	monkeys := []Monkey{}

	lineIdx := 0
	for lineIdx < len(lines) {
		monkey := Monkey{}
		name := lines[lineIdx]
		monkey.Id, _ = strconv.Atoi(strings.Split(name, " ")[1][0:1])
		items := strings.Split(lines[lineIdx+1], ": ")[1]
		monkey.Items = StrsToIns(strings.Split(items, ", "))
		operation := strings.Split(strings.Split(lines[lineIdx+2], "= old ")[1], " ")
		if operation[0] == "+" {
			monkey.PlusOrMult = true
		} else {
			monkey.PlusOrMult = false
		}
		if operation[1] == "old" {
			monkey.Same = true
		} else {
			by, _ := strconv.Atoi(operation[1])
			monkey.By = by
		}
		monkey.Test, _ = strconv.Atoi(strings.Split(lines[lineIdx+3], "by ")[1])
		monkey.TestTrue, _ = strconv.Atoi(strings.Split(lines[lineIdx+4], "monkey ")[1])
		monkey.TestFalse, _ = strconv.Atoi(strings.Split(lines[lineIdx+5], "monkey ")[1])

		monkeys = append(monkeys, monkey)
		lineIdx += 7
	}

	round := 0
	for round < 20 {
		for _, mon := range monkeys {
			for _, item := range mon.Items {
				worry := item

				// Operation
				if mon.PlusOrMult {
					if mon.Same {
						worry += worry
					} else {
						worry += mon.By
					}
				} else {
					if mon.Same {
						worry *= worry
					} else {
						worry *= mon.By
					}
				}

				// Boredom
				worry /= 3

				// Test
				if worry%mon.Test == 0 {
					monkeys[mon.TestTrue].Items = append(monkeys[mon.TestTrue].Items, worry)
				} else {
					monkeys[mon.TestFalse].Items = append(monkeys[mon.TestFalse].Items, worry)
				}

			}
			monkeys[mon.Id].Inspected += len(mon.Items)
			monkeys[mon.Id].Items = []int{}
		}
		round++
	}

	// Calculation
	totals := []int{}
	for _, mon := range monkeys {
		totals = append(totals, mon.Inspected)
	}
	sort.Ints(totals)
	active1 := totals[len(totals)-1]
	active2 := totals[len(totals)-2]

	fmt.Println(active1 * active2)
}
