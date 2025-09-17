package advent

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220401() {
	lines := advent.GetLinesFromFile("advent/2022/04/ie.txt")
	// lines := advent.GetLinesFromFile("advent/2022/04/input.txt")
	total := 0
	for _, v := range lines {
		elves := strings.Split(v, ",")
		cond1 := true
		cond2 := true
		e1Sections := strings.Split(elves[0], "-")
		e2Sections := strings.Split(elves[1], "-")
		e1a, _ := strconv.Atoi(e1Sections[0])
		e1b, _ := strconv.Atoi(e1Sections[1])
		e2a, _ := strconv.Atoi(e2Sections[0])
		e2b, _ := strconv.Atoi(e2Sections[1])
		e1 := []int{}
		e2 := []int{}
		for i := e1a; i <= e1b; i++ {
			e1 = append(e1, i)
		}
		for i := e2a; i <= e2b; i++ {
			e2 = append(e2, i)
		}
		for _, v1 := range e2 {
			if !slices.Contains(e1, v1) {
				cond1 = false
				break
			}
		}
		for _, v2 := range e1 {
			if !slices.Contains(e2, v2) {
				cond2 = false
				break
			}
		}
		if cond1 || cond2 {
			total++
		}
	}
	fmt.Println(total)
}
