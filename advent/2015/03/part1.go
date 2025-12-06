package advent

import (
	"fmt"

	"github.com/kengru/problems/advent"
)

func Year20150301() {
	characters := advent.GetCharactersFromFile("advent/2015/03/input.txt")

	x, y := 0, 0
	set := make(map[string]struct{})
	firstPosition := fmt.Sprintf("%v,%v", x, y)
	set[firstPosition] = struct{}{}

	for _, c := range characters {
		switch c {
		case ">":
			x++
		case "<":
			x--
		case "^":
			y++
		case "v":
			y--
		}

		position := fmt.Sprintf("%v,%v", x, y)
		set[position] = struct{}{}
	}

	fmt.Println(len(set))
}
