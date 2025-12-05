package advent

import (
	"fmt"

	"github.com/kengru/problems/advent"
)

func Year20150101() {
	characters := advent.GetCharactersFromFile("advent/2015/01/input.txt")

	floor := 0
	for _, ch := range characters {
		if ch == "(" {
			floor++
		} else {
			floor--
		}
	}

	fmt.Println(floor)
}
