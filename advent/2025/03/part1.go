package advent

import (
	"fmt"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20250301() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2025/03/ie.txt")
	lines := advent.GetLinesFromFile("advent/2025/03/input.txt")

	joltage := 0
	lineIdx := 0
	for lineIdx < len(lines) {
		bank := lines[lineIdx]
		fmt.Println(bank)
		beg := 0
		end := len(bank)
		bank_joltage := 0
		for i := beg; i < end-1; i++ {
			for j := beg + 1; j < end; j++ {
				jolt, _ := strconv.Atoi(fmt.Sprintf("%c%c", bank[i], bank[j]))
				if jolt > bank_joltage {
					bank_joltage = jolt
				}
			}
			beg++
		}

		fmt.Println(bank_joltage)
		joltage += bank_joltage

		lineIdx++
	}

	fmt.Println(joltage)
}

// // Getting input
// lines := advent.GetLinesFromFile("advent/2025/03/ie.txt")
// // line := advent.GetLineFromFile("advent/2025/03/input.txt")

// joltage := 0
// lineIdx := 0
// for lineIdx < len(lines) {
// 	bank := lines[lineIdx]
// 	fmt.Println(bank)
// 	current := 9
// 	from := 0
// 	var j bytes.Buffer
// 	for current >= 0 || len(j.String()) < 2 {
// 		found := false
// 		for i := from; i < len(bank); i++ {
// 			if bank[i] == byte(current) {
// 				j.WriteString(string(bank[i]))
// 				found = true
// 				from = i + 1
// 				current = 9
// 				fmt.Println(j.String())
// 			}
// 		}
// 		if !found {
// 			current--
// 		}
// 	}

// 	bank_joltage, _ := strconv.Atoi(j.String())
// 	joltage += bank_joltage

// 	lineIdx++
// }

// fmt.Println(joltage)
