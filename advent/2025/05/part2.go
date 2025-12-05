package advent

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20250502() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2025/05/ie.txt")
	lines := advent.GetLinesFromFile("advent/2025/05/input.txt")

	lineIdx := 0
	rngs := make([]Range, 0)
	fresh := 0

	for lineIdx < len(lines) {
		rng := lines[lineIdx]
		if rng == "" {
			break
		}
		spl := strings.Split(rng, "-")
		min, _ := strconv.Atoi(spl[0])
		max, _ := strconv.Atoi(spl[1])
		rngs = append(rngs, Range{
			min, max,
		})

		lineIdx++
	}

	mergedIdx := 1
	merged := make([]*Range, 0)
	merged = append(merged, &rngs[0])
	slices.SortFunc(rngs, func(a, b Range) int {
		return a.min - b.min
	})

	for mergedIdx < len(rngs) {
		check := rngs[mergedIdx]
		mergedOr := false
		for _, m := range merged {
			if check.min <= m.max {
				m.min = int(math.Min(float64(m.min), float64(check.min)))
				m.max = int(math.Max(float64(m.max), float64(check.max)))
				mergedOr = true
				break
			}
		}
		if !mergedOr {
			merged = append(merged, &check)
		}
		mergedIdx++
	}

	for _, m := range merged {
		fresh += m.max - m.min + 1
	}
	fmt.Println(fresh)
}
