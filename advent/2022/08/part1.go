package advent

import (
	"fmt"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20220801() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2022/08/ie.txt")
	lines := advent.GetLinesFromFile("advent/2022/08/input.txt")

	visible, columns, rows := 0, 0, 0
	grid := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, r := range line {
			value, _ := strconv.Atoi(string(r))
			row = append(row, value)
		}
		rows = len(row)
		grid = append(grid, row)
	}
	columns = len(lines)
	visible = (rows * 2) + ((columns - 2) * 2)

	for j := 1; j < len(grid)-1; j++ {
		for i := 1; i < len(grid[j])-1; i++ {
			tree := grid[j][i]
			tcheck, bcheck, lcheck, rcheck := true, true, true, true

			// top check
			for tp := j - 1; tp >= 0; tp-- {
				if grid[tp][i] >= tree {
					tcheck = false
					break
				}
			}
			// bot check
			for bt := j + 1; bt < len(grid[j]); bt++ {
				if grid[bt][i] >= tree {
					bcheck = false
					break
				}
			}
			// left check
			for lf := i - 1; lf >= 0; lf-- {
				if grid[j][lf] >= tree {
					lcheck = false
					break
				}
			}
			// right check
			for rg := i + 1; rg < len(grid[i]); rg++ {
				if grid[j][rg] >= tree {
					rcheck = false
					break
				}
			}

			if tcheck || bcheck || lcheck || rcheck {
				visible++
			}
		}
	}

	fmt.Println(visible)
}
