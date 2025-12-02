package advent

import (
	"fmt"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20250101() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2025/01/ie.txt")
	lines := advent.GetLinesFromFile("advent/2025/01/input.txt")

	starting := 50
	zero := 0

	lineIdx := 0
	for lineIdx < len(lines) {
		dir := lines[lineIdx][0:1]
		amt, _ := strconv.Atoi(lines[lineIdx][1:])
		if dir == "R" {
			starting += amt
			if starting >= 100 {
				starting = starting % 100
			}
		} else {
			starting -= amt
			if starting < 0 {
				starting = 100 + (starting % 100)
			}
		}
		if starting == 100 {
			starting = 0
		}

		if starting == 0 {
			zero++
		}
		fmt.Println(starting)

		lineIdx++
	}

	fmt.Println(zero)
}
