package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20150201() {
	lines := advent.GetLinesFromFile("advent/2015/02/input.txt")

	value := 0
	for _, line := range lines {
		dims := strings.Split(line, "x")
		x, _ := strconv.Atoi(dims[0])
		y, _ := strconv.Atoi(dims[1])
		z, _ := strconv.Atoi(dims[2])
		v1 := x * y
		v2 := x * z
		v3 := y * z
		sfa := 2*v1 + 2*v2 + 2*v3
		value += sfa + min(v1, v2, v3)
	}

	fmt.Println(value)
}
