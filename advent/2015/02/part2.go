package advent

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20150202() {
	lines := advent.GetLinesFromFile("advent/2015/02/input.txt")

	value := 0
	for _, line := range lines {
		dims := strings.Split(line, "x")
		slices.SortFunc(dims, func(a, b string) int {
			astr, _ := strconv.Atoi(a)
			bstr, _ := strconv.Atoi(b)
			return astr - bstr
		})
		x, _ := strconv.Atoi(dims[0])
		y, _ := strconv.Atoi(dims[1])
		z, _ := strconv.Atoi(dims[2])
		rw := 2*x + 2*y
		bow := x * y * z
		value += rw + bow
	}

	fmt.Println(value)
}
