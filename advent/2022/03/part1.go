package advent

import (
	"fmt"
	"unicode"

	"github.com/kengru/problems/advent"
)

func Year20220301() {
	// lines := advent.GetLinesFromFile("advent/2022/03/ie.txt")
	lines := advent.GetLinesFromFile("advent/2022/03/input.txt")
	total := 0
	for _, v := range lines {
		fmp := make(map[rune]bool)
		smp := make(map[rune]bool)
		for i, r := range v {
			if i < len(v)/2 {
				_, ex := fmp[r]
				if !ex {
					fmp[r] = true
				}
			} else {
				_, ex := smp[r]
				if !ex {
					smp[r] = true
				}
			}
		}
		r := 'a'
		for k := range fmp {
			for k2 := range smp {
				if k == k2 {
					r = k
				}
			}
		}
		fmt.Println("common: ", string(r), valueRune(r))
		total += valueRune(r)
	}
	fmt.Println(total)
}

func valueRune(r rune) int {
	if unicode.IsLower(r) {
		output1 := r - 'a'
		output2 := (26 - 1) / ('z' - 'a')
		return int(output1*output2 + 1)
	}
	output1 := r - 'A'
	output2 := (52 - 27) / ('Z' - 'A')
	return int(output1*output2 + 27)
}
