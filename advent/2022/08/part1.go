package advent

import (
	"fmt"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20220801() {
	// Getting input
	lines := advent.GetLinesFromFile("advent/2022/08/ie.txt")
	// lines := advent.GetLinesFromFile("advent/2022/08/input.txt")

	grid := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, r := range line {
			value, _ := strconv.Atoi(string(r))
			row = append(row, value)
		}
		grid = append(grid, row)
	}
	fmt.Println(grid)
}
