package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20250201() {
	// Getting input
	// line := advent.GetLineFromFile("advent/2025/02/ie.txt")
	line := advent.GetLineFromFile("advent/2025/02/input.txt")

	invalids := 0
	ranges := strings.Split(line, ",")
	for _, rng := range ranges {
		splt := strings.Split(rng, "-")
		lower, _ := strconv.Atoi(splt[0])
		higher, _ := strconv.Atoi(splt[1])
		for i := lower; i <= higher; i++ {
			st := strconv.Itoa(i)
			if isInvalid(st) {
				invalids += i
			}
		}
	}

	fmt.Println(invalids)
}

func isInvalid(id string) bool {
	length := len(id)
	if length%2 != 0 {
		return false
	}
	first, _ := strconv.Atoi(id[0 : length/2])
	second, _ := strconv.Atoi(id[length/2:])
	fmt.Println(first, second)
	return first == second
}
