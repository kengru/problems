package advent

import (
	"fmt"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220202() {
	lines := advent.GetLinesFromFile("advent/2022/02/input.txt")
	mapa := make(map[string]int)
	w := make(map[string]string)
	l := make(map[string]string)
	mapa["A"] = 1
	mapa["B"] = 2
	mapa["C"] = 3
	w["A"] = "B"
	w["B"] = "C"
	w["C"] = "A"
	l["A"] = "C"
	l["B"] = "A"
	l["C"] = "B"
	total := 0
	for _, v := range lines {
		t := 0
		hands := strings.Split(v, " ")
		if hands[1] == "Y" {
			t += 3 + mapa[hands[0]]
		}
		if hands[1] == "X" {
			t += 0 + mapa[l[hands[0]]]
		}
		if hands[1] == "Z" {
			t += 6 + mapa[w[hands[0]]]
		}
		total += t + mapa[hands[1]]
	}
	fmt.Println(total)
}
