package advent

import (
	"fmt"

	"github.com/kengru/problems/advent"
)

func Year20150302() {
	characters := advent.GetCharactersFromFile("advent/2015/03/input.txt")

	x, y, x2, y2 := 0, 0, 0, 0
	toggle := false

	set := make(map[string]struct{})
	firstPosition := fmt.Sprintf("%v,%v", x, y)
	set[firstPosition] = struct{}{}

	for _, c := range characters {
		if toggle {
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
		} else {
			switch c {
			case ">":
				x2++
			case "<":
				x2--
			case "^":
				y2++
			case "v":
				y2--
			}
		}

		toggle = !toggle
		position := ""
		if !toggle {
			position = fmt.Sprintf("%v,%v", x, y)
		} else {
			position = fmt.Sprintf("%v,%v", x2, y2)
		}
		set[position] = struct{}{}
	}

	fmt.Println(len(set))
}
