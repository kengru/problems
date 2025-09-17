package advent

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20220101() {
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
	fmt.Println(elves[len(elves)-1])
}
