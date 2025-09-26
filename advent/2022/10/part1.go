package advent

import (
	"fmt"
	"math"
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
		y := 0.5 - 0.5*math.Cos(float64(math.Pi/20.0)*float64(cycle))
		if y == 1 {
			fmt.Println("here", cycle, registerX)
			signal += registerX * cycle
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
