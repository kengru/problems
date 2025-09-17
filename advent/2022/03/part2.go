package advent

import (
	"fmt"

	"github.com/kengru/problems/advent"
)

func Year20220302() {
	// lines := advent.GetLinesFromFile("advent/2022/03/ie.txt")
	lines := advent.GetLinesFromFile("advent/2022/03/input.txt")
	total := 0
	for i := 0; i < len(lines); i += 3 {
		fmp := make(map[rune]bool)
		smp := make(map[rune]bool)
		tmp := make(map[rune]bool)
		for _, r := range lines[i] {
			_, ex := fmp[r]
			if !ex {
				fmp[r] = true
			}
		}
		for _, r := range lines[i+1] {
			_, ex := smp[r]
			if !ex {
				smp[r] = true
			}
		}
		for _, r := range lines[i+2] {
			_, ex := tmp[r]
			if !ex {
				tmp[r] = true
			}
		}
		r := 'a'
		k12 := make(map[rune]bool)
		for k := range fmp {
			for k2 := range smp {
				if k == k2 {
					_, ex := k12[k]
					if !ex {
						k12[k] = true
					}
				}
			}
		}
		for k := range k12 {
			for k2 := range tmp {
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
