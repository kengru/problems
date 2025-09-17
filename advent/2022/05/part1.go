package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220501() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2022/05/ie.txt")
	lines := advent.GetLinesFromFile("advent/2022/05/input.txt")

	header := []string{}
	instructions := []string{}

	// Creating header
	idx := 0
	for idx < len(lines)-1 {
		if lines[idx] == "" {
			idx++
			break
		}
		header = append(header, lines[idx])
		idx++
	}

	// Creating instructions
	for idx < len(lines) {
		instructions = append(instructions, lines[idx])
		idx++
	}

	// Amount of stacks to be created
	amount := 0
	idx = 1
	for idx < len(header[len(header)-1])-1 {
		amount++
		idx += 4
	}

	// Filling stacks
	idx = len(header) - 2
	stacks := make([]advent.Stack[string], amount)
	for idx >= 0 {
		line := header[idx]
		for i := 1; i < len(line); i += 4 {
			letter := line[i]
			if string(letter) != " " {
				stacks[(i-1)/4].Add(string(letter))
			}
		}
		idx--
	}

	// Executing instructions
	for _, ins := range instructions {
		amount, from, to := interpreterInstructions(ins)
		temp := []string{}
		for range amount {
			temp = append(temp, stacks[from-1].Pop())
		}
		for i := range len(temp) {
			stacks[to-1].Add(temp[i])
		}
		// for i := len(temp) - 1; i >= 0; i-- {
		// 	stacks[to-1].Add(temp[i])
		// }
	}

	crates := ""

	for _, st := range stacks {
		crates += st.Peek()
	}

	fmt.Println(crates)
}

// Getting values from instructions
func interpreterInstructions(ins string) (amount int, from int, to int) {
	splt := strings.Split(ins, " ")
	amount, _ = strconv.Atoi(splt[1])
	from, _ = strconv.Atoi(splt[3])
	to, _ = strconv.Atoi(splt[5])
	return
}
