package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

type Coord struct {
	X int
	Y int
}

func Year20220902() {
	// Getting input
	lines := advent.GetLinesFromFile("advent/2022/09/ie2.txt")
	// lines := advent.GetLinesFromFile("advent/2022/09/input.txt")

	rope := [10]Coord{}
	visited := make(map[string]bool)

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		dir := instruction[0]
		amt, _ := strconv.Atoi(instruction[1])

		for amt > 0 {
			if dir == "R" {
				rope[0].X++
			}
			if dir == "D" {
				rope[0].Y--
			}
			if dir == "L" {
				rope[0].X--
			}
			if dir == "U" {
				rope[0].Y++
			}
			for i := range 9 {
				current := &rope[i]
				next := &rope[i+1]

				if dir == "R" {
					diffX := current.X - next.X
					if diffX > 1 {
						next.X++
						diffY := current.Y - next.Y
						if diffY != 0 {
							next.Y += diffY
						}
					}
				}
				if dir == "D" {
					diffY := next.Y - current.Y
					if diffY > 1 {
						next.Y--
						diffX := current.X - next.X
						if diffX != 0 {
							next.X += diffX
						}
					}
				}
				if dir == "L" {
					diffX := next.X - current.X
					if diffX > 1 {
						next.X--
						diffY := current.Y - next.Y
						if diffY != 0 {
							next.Y += diffY
						}
					}
				}
				if dir == "U" {
					diffY := current.Y - next.Y
					if diffY > 1 {
						next.Y++
						diffX := current.X - next.X
						if diffX != 0 {
							next.X += diffX
						}
					}
				}
			}
			tailLoc := fmt.Sprintf("%d,%d", rope[9].X, rope[9].Y)
			visited[tailLoc] = true
			amt--
		}

	}

	fmt.Println(len(visited))
}
