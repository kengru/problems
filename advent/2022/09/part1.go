package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220901() {
	// Getting input
	lines := advent.GetLinesFromFile("advent/2022/09/ie.txt")
	// lines := advent.GetLinesFromFile("advent/2022/09/input.txt")

	hx, hy := 0, 0
	tx, ty := 0, 0
	visited := make(map[string]bool)

	for _, line := range lines {
		instruction := strings.Split(line, " ")
		dir := instruction[0]
		amt, _ := strconv.Atoi(instruction[1])
		for amt > 0 {
			amt--
		}
		diffY := hy - ty
		diffX := hx - tx
	}

	fmt.Println(len(visited))
}
