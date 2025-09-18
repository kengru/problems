package advent

import (
	"fmt"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220601() {
	// Getting input
	// chars := advent.GetCharactersFromFile("advent/2022/06/ie.txt")
	chars := advent.GetCharactersFromFile("advent/2022/06/input.txt")
	i1 := 0
	i2 := 3
	for i2 < len(chars) {
		four := strings.Join(chars[i1:i2+1], "")
		if uniques(four) {
			break
		}
		i1++
		i2++
	}

	fmt.Println(i2 + 1)
}

func uniques(fu string) bool {
	mapa := make(map[rune]bool)
	for _, vv := range fu {
		_, ex := mapa[vv]
		if ex {
			return false
		} else {
			mapa[vv] = true
		}
	}
	return true
}
