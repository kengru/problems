package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20221002() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2022/10/ie2.txt")
	lines := advent.GetLinesFromFile("advent/2022/10/input.txt")

	registerX := 1
	crt := make([][]string, 6)
	x, y := 0, 0
	cycle := 1
	lineIdx := 0
	shouldAdd := false
	for lineIdx < len(lines) {
		if x == 0 {
			crt[y] = make([]string, 40)
		}
		line := lines[lineIdx]
		if line == "noop" {
			cycle++
			lineIdx++
			if x >= registerX-1 && x <= registerX+1 {
				crt[y][x] = "#"
			} else {
				crt[y][x] = "."
			}
			x++
			if x > 39 {
				y++
				x = 0
			}
			continue
		}
		instruction := strings.Split(line, " ")
		amt, _ := strconv.Atoi(instruction[1])
		if x >= registerX-1 && x <= registerX+1 {
			crt[y][x] = "#"
		} else {
			crt[y][x] = "."
		}
		if shouldAdd {
			registerX += amt
			lineIdx++
			shouldAdd = false
		} else {
			shouldAdd = true
		}
		x++
		if x > 39 {
			y++
			x = 0
		}
		cycle++
	}

	for _, v := range crt {
		for _, u := range v {
			fmt.Print(u)
		}
		fmt.Println()
	}
}
