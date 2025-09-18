package advent

import (
	"fmt"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220602() {
	// Getting input
	// chars := advent.GetCharactersFromFile("advent/2022/06/ie.txt")
	chars := advent.GetCharactersFromFile("advent/2022/06/input.txt")
	i1 := 0
	i2 := 13
	for i2 < len(chars) {
		fourteen := strings.Join(chars[i1:i2+1], "")
		if uniques(fourteen) {
			break
		}
		i1++
		i2++
	}

	fmt.Println(i2 + 1)
}
