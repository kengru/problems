package advent

import (
	"fmt"
	"math"

	"github.com/kengru/problems/advent"
)

func Year20250402() {
	// Getting input
	// lines := advent.GetMatrixFromFile("advent/2025/04/ie.txt")
	lines := advent.GetMatrixFromFile("advent/2025/04/input.txt")

	total := 0
	value := 0
	firstTime := true
	for value != 0 || firstTime {
		firstTime = false
		value, lines = totalCount(&lines)
		total += value
	}

	fmt.Println(total)
}

func countArround2(matrix [][]string, x int, y int) int {
	count := 0
	sizeY := len(matrix)
	sizeX := len(matrix[0])
	for i := int(math.Max(0, (float64(y - 1)))); i <= int(math.Min(float64(sizeY-1), float64(y+1))); i++ {
		for j := int(math.Max(0, (float64(x - 1)))); j <= int(math.Min(float64(sizeX-1), float64(x+1))); j++ {
			if string(matrix[i][j]) == "@" {
				count++
			}
		}
	}
	return count
}

func totalCount(lines *[][]string) (int, [][]string) {
	total := 0
	sizeY := len(*lines)
	sizeX := len((*lines)[0])
	newLines := make([][]string, 0)
	for i := 0; i < sizeY; i++ {
		line := make([]string, 0)
		for j := 0; j < sizeX; j++ {
			current := (*lines)[i][j]
			if string(current) == "@" {
				count := countArround2(*lines, j, i) - 1
				if count < 4 {
					line = append(line, ".")
					total++
				} else {
					line = append(line, "@")
				}
			} else {
				line = append(line, ".")
			}
		}
		newLines = append(newLines, line)
	}
	return total, newLines
}
