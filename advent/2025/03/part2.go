package advent

import (
	"fmt"
	"math"
	"strconv"

	"github.com/kengru/problems/advent"
)

func Year20250302() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2025/03/ie.txt")
	lines := advent.GetLinesFromFile("advent/2025/03/input.txt")

	joltage := 0
	for _, bank := range lines {
		finalLength := 12
		bankLength := len(bank)

		left := 0
		value := 0
		for finalLength > 0 {
			winnerIdx := math.MaxInt
			winnerNumber := 0
			for left < bankLength-finalLength+1 {
				number, _ := strconv.Atoi(string(bank[left]))
				if number > winnerNumber && number != winnerNumber {
					winnerIdx = left
					winnerNumber = number
				}
				left++
			}
			finalLength--
			value += winnerNumber * int(math.Pow10(finalLength))
			left = winnerIdx + 1
		}

		joltage += value
	}

	fmt.Println(joltage)
}
