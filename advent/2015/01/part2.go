package advent

import (
	"fmt"

	"github.com/kengru/problems/advent"
)

func Year20150102() {
	characters := advent.GetCharactersFromFile("advent/2015/01/input.txt")

	floor := 0
	for idx, ch := range characters {
		if floor == -1 {
			floor = idx
			break
		}
		if ch == "(" {
			floor++
		} else {
			floor--
		}
	}

	fmt.Println(floor)
}
