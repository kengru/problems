package advent

import (
	"fmt"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220201() {
	lines := advent.GetLinesFromFile("advent/2022/02/input.txt")
	mapa := make(map[string]int)
	mapa["A"] = 1
	mapa["B"] = 2
	mapa["C"] = 3
	mapa["X"] = 1
	mapa["Y"] = 2
	mapa["Z"] = 3
	total := 0
	for _, v := range lines {
		t := 0
		hands := strings.Split(v, " ")
		if hands[1] == "X" && hands[0] == "C" {
			t += 6
		}
		if hands[1] == "Y" && hands[0] == "A" {
			t += 6
		}
		if hands[1] == "Z" && hands[0] == "B" {
			t += 6
		}
		if mapa[hands[0]] == mapa[hands[1]] {
			t += 3
		}
		total += t + mapa[hands[1]]
	}
	fmt.Println(total)
}
