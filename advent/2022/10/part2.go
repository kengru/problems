package advent

import (
	"fmt"
	"math"
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
	// lines := advent.GetLinesFromFile("advent/2022/09/ie.txt")
	// lines := advent.GetLinesFromFile("advent/2022/09/ie2.txt")
	lines := advent.GetLinesFromFile("advent/2022/09/input.txt")

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

				dX := current.X - next.X
				dY := current.Y - next.Y

				diffX := int(math.Abs(float64(dX)))
				diffY := int(math.Abs(float64(dY)))
				if diffX > 1 {
					if dX > 1 {
						next.X += dX - 1
					} else {
						next.X += dX + 1
					}
					if dY != 0 {
						if dY > 1 {
							next.Y += dY - 1
						} else if dY < -1 {
							next.Y += dY + 1
						} else {
							next.Y += dY
						}
					}
					continue
				}
				if diffY > 1 {
					if dY > 1 {
						next.Y += dY - 1
					} else {
						next.Y += dY + 1
					}
					if dX != 0 {
						if dX > 1 {
							next.X += dX - 1
						} else if dX < -1 {
							next.X += dX + 1
						} else {
							next.X += dX
						}
					}
					continue
				}
			}
			tailLoc := fmt.Sprintf("%d,%d", rope[9].X, rope[9].Y)
			visited[tailLoc] = true
			amt--
		}

	}

	fmt.Println(len(visited))
}
