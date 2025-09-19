package advent

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20220802() {
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

	routes := []int{}
	for j := 1; j < len(grid)-1; j++ {
		for i := 1; i < len(grid[j])-1; i++ {
			tree := grid[j][i]
			scenic := 1

			// top check
			ts := 1
			for tp := j - 1; tp >= 0; tp-- {
				if grid[tp][i] >= tree {
					scenic *= ts
					break
				}
				ts++
			}
			// bot check
			bs := 1
			for bt := j + 1; bt < len(grid[j]); bt++ {
				if grid[bt][i] >= tree {
					scenic *= bs
					break
				}
				bs++
			}
			// left check
			ls := 1
			for lf := i - 1; lf >= 0; lf-- {
				if grid[j][lf] >= tree {
					scenic *= ls
					break
				}
				ls++
			}
			// right check
			rs := 1
			for rg := i + 1; rg < len(grid[i]); rg++ {
				if grid[j][rg] >= tree {
					scenic *= rs
					break
				}
				rs++
			}

			routes = append(routes, scenic)
		}
	}

	sort.Ints(routes)
	fmt.Println(routes)
}
