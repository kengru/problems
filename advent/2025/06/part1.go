package advent

import (
	"fmt"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20250601() {
	// lines := advent.GetSeparatedFromFile("advent/2025/06/ie.txt")
	lines := advent.GetSeparatedFromFile("advent/2025/06/input.txt", true)

	idx := 0
	wide := len(lines[0])
	total := 0
	for idx < wide {
		calc := 0
		numbers := make([]int, 0)
		for i := 0; i < len(lines); i++ {
			value := lines[i][idx]
			if i < len(lines)-1 {
				nv, _ := strconv.Atoi(value)
				numbers = append(numbers, nv)
			} else {
				switch value {
				case "+":
					for _, n := range numbers {
						calc += n
					}
				case "*":
					calc = 1
					for _, n := range numbers {
						calc *= n
					}
				}
			}
		}
		idx++
		total += calc
	}

	fmt.Println(total)
}
