package advent

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20240102() {
	lines := advent.GetLinesFromFile("advent/2024/01/input.txt")
	list1 := make([]int, 0)
	list2 := make([]int, 0)
	similarity := 0

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
		times := 0
		for _, b := range list2 {
			if b == list1[i] {
				times++
			}
		}
		similarity += list1[i] * times
	}

	fmt.Println(similarity)
}
