package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220901() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2022/09/ie.txt")
	lines := advent.GetLinesFromFile("advent/2022/09/input.txt")

	hx, hy := 0, 0
	tx, ty := 0, 0
	visited := make(map[string]bool)

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		dir := instruction[0]
		amt, _ := strconv.Atoi(instruction[1])
		for amt > 0 {
			if dir == "R" {
				hx++
				diffX := hx - tx
				if diffX > 1 {
					tx++
					diffY := hy - ty
					if diffY != 0 {
						ty += diffY
					}
				}
			}
			if dir == "D" {
				hy--
				diffY := ty - hy
				if diffY > 1 {
					ty--
					diffX := hx - tx
					if diffX != 0 {
						tx += diffX
					}
				}
			}
			if dir == "L" {
				hx--
				diffX := tx - hx
				if diffX > 1 {
					tx--
					diffY := hy - ty
					if diffY != 0 {
						ty += diffY
					}
				}
			}
			if dir == "U" {
				hy++
				diffY := hy - ty
				if diffY > 1 {
					ty++
				}
				diffX := hx - tx
				if diffX != 0 {
					tx += diffX
				}
			}

			tailLoc := fmt.Sprintf("%d,%d", tx, ty)
			visited[tailLoc] = true
			amt--
		}
	}

	fmt.Println(len(visited))
}
