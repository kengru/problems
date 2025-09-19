package advent

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20220802() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2022/08/ie.txt")
	lines := advent.GetLinesFromFile("advent/2022/08/input.txt")

	grid := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, r := range line {
			value, _ := strconv.Atoi(string(r))
			row = append(row, value)
		}
		grid = append(grid, row)
	}

	routes := []int{}
	for j := 1; j < len(grid)-1; j++ {
		for i := 1; i < len(grid[j])-1; i++ {
			tree := grid[j][i]

			// top check
			ts := 0
			for tp := j - 1; tp >= 0; tp-- {
				if grid[tp][i] >= tree {
					ts++
					break
				}
				ts++
			}
			// bot check
			bs := 0
			for bt := j + 1; bt < len(grid[j]); bt++ {
				if grid[bt][i] >= tree {
					bs++
					break
				}
				bs++
			}
			// left check
			ls := 0
			for lf := i - 1; lf >= 0; lf-- {
				if grid[j][lf] >= tree {
					ls++
					break
				}
				ls++
			}
			// right check
			rs := 0
			for rg := i + 1; rg < len(grid[i]); rg++ {
				if grid[j][rg] >= tree {
					rs++
					break
				}
				rs++
			}

			routes = append(routes, ts*bs*ls*rs)
		}
	}

	sort.Ints(routes)
	fmt.Println(routes[len(routes)-1])
}
