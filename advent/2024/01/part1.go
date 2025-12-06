package advent

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20240101() {
	lines := advent.GetLinesFromFile("advent/2024/01/input.txt")
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	dist := 0

	for _, line := range lines {
		both := strings.Split(line, "   ")
		a, _ := strconv.Atoi(both[0])
		b, _ := strconv.Atoi(both[1])
		list1 = append(list1, a)
		list2 = append(list2, b)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	for i := range lines {
		a1 := float64(list1[i])
		a2 := float64(list2[i])
		dist += int(math.Abs(a1 - a2))
	}

	fmt.Println(dist)
}
