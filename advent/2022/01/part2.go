package advent

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20220102() {
	lines := advent.GetLinesFromFile("advent/2022/01/input.txt")
	elves := []int{}
	total := 0
	for _, v := range lines {
		if v == "" {
			elves = append(elves, total)
			total = 0
		} else {
			value, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			total += value
		}
	}
	slices.Sort(elves)
	top := len(elves) - 1
	total = elves[top] + elves[top-1] + elves[top-2]
	fmt.Println(total)
}
