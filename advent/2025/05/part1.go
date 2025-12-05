package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

type Range struct {
	min int
	max int
}

func Year20250501() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2025/05/ie.txt")
	lines := advent.GetLinesFromFile("advent/2025/05/input.txt")

	lineIdx := 0
	fresh := 0
	ranges := make([]Range, 0)

	for lineIdx < len(lines) {
		rng := lines[lineIdx]
		if rng == "" {
			break
		}
		spl := strings.Split(rng, "-")
		min, _ := strconv.Atoi(spl[0])
		max, _ := strconv.Atoi(spl[1])
		ranges = append(ranges, Range{
			min: min,
			max: max,
		})
		lineIdx++
	}

	for lineIdx < len(lines) {
		ingredient, _ := strconv.Atoi(lines[lineIdx])
		for _, r := range ranges {
			if ingredient >= r.min && ingredient <= r.max {
				fresh++
				break
			}
		}
		lineIdx++
	}

	fmt.Println(fresh)
}
