package advent

import (
	"fmt"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20250102() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2025/01/ie.txt") // 6
	lines := advent.GetLinesFromFile("advent/2025/01/input.txt") // 2261, 5067, 6007, 5899

	position := 50
	zero := 0

	lineIdx := 0
	for lineIdx < len(lines) {
		var dirs int
		direction := lines[lineIdx][0:1]
		distance, _ := strconv.Atoi(lines[lineIdx][1:])
		fmt.Println("dir:", direction, "amt:", distance, "start:", position, "zero:", zero)

		previous := position

		if direction == "R" {
			dirs = 1
		} else {
			dirs = -1
		}

		if distance >= 100 {
			zero += distance / 100
		}

		position += (distance % 100) * dirs

		if position > 99 {
			position -= 100
			if position != 0 && previous != 0 {
				zero++
			}
		}
		if position < 0 {
			position += 100
			if position != 0 && previous != 0 {
				zero++
			}
		}
		if position == 0 {
			zero++
		}
		fmt.Println(position)

		lineIdx++
	}

	fmt.Println("total:", zero)
}
