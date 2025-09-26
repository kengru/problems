package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20221001() {
	// Getting input
	lines := advent.GetLinesFromFile("advent/2022/10/ie2.txt")
	// lines := advent.GetLinesFromFile("advent/2022/09/input.txt")

	cycle := 0
	signal := 0
	registerX := 1
	lineIdx := 0
	shouldAdd := false
	for lineIdx < len(lines) {
		line := lines[lineIdx]
		if cycle == 20 {
			println("at 20", registerX)
			signal += registerX * 20
		}
		if cycle == 60 {
			println("at 60", registerX)
			signal += registerX * 60
		}
		if line == "noop" {
			cycle++
			lineIdx++
			continue
		}
		instruction := strings.Split(line, " ")
		amt, _ := strconv.Atoi(instruction[1])
		if shouldAdd {
			registerX += amt
			lineIdx++
			shouldAdd = false
		} else {
			shouldAdd = true
		}
		cycle++
	}

	fmt.Println("cycle", cycle)
	fmt.Println("signal", signal)
	fmt.Println("registerX", registerX)
}
