package advent

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20221102() {
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

	lcm := 1
	for _, mon := range monkeys {
		lcm *= mon.Test
	}

	round := 0
	for round < 10000 {
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

				// Test
				if worry%mon.Test == 0 {
					monkeys[mon.TestTrue].Items = append(monkeys[mon.TestTrue].Items, worry%lcm)
				} else {
					monkeys[mon.TestFalse].Items = append(monkeys[mon.TestFalse].Items, worry%lcm)
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
