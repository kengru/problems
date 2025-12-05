package advent

import (
	"fmt"
	"math"

	"github.com/kengru/problems/advent"
)

func Year20250401() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2025/04/ie.txt")
	lines := advent.GetLinesFromFile("advent/2025/04/input.txt")

	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			current := lines[i][j]
			if string(current) == "@" {
				count := countArround(lines, j, i) - 1
				if count < 4 {
					total++
				}
			}
		}
	}

	fmt.Println(total)
}

func countArround(matrix []string, x int, y int) int {
	count := 0
	for i := int(math.Max(0, (float64(y - 1)))); i <= int(math.Min(float64(len(matrix)-1), float64(y+1))); i++ {
		for j := int(math.Max(0, (float64(x - 1)))); j <= int(math.Min(float64(len(matrix[i])-1), float64(x+1))); j++ {
			if string(matrix[i][j]) == "@" {
				count++
			}
		}
	}
	return count
}
