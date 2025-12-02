package advent

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20250202() {
	// Getting input
	// line := advent.GetLineFromFile("advent/2025/02/ie.txt")
	line := advent.GetLineFromFile("advent/2025/02/input.txt")

	invalids := 0
	ranges := strings.Split(line, ",")
	for _, rng := range ranges {
		splt := strings.Split(rng, "-")
		lower, _ := strconv.Atoi(splt[0])
		higher, _ := strconv.Atoi(splt[1])
		for i := lower; i <= higher; i++ {
			st := strconv.Itoa(i)
			if isNewInvalid(st) {
				invalids += i
			}
		}
	}

	fmt.Println(invalids)
}

func isNewInvalid(id string) bool {
	length := len(id)
	invalid := false
	for i := 1; i <= length/2; i++ {
		subs := goodSplit(id, i)
		if len(subs) == 1 {
			break
		}
		eq := true
		for j := 0; j < len(subs)-1; j++ {
			if subs[j] != subs[j+1] {
				eq = false
			}
		}
		if eq {
			return true
		}
	}
	return invalid
}

func goodSplit(id string, n int) []string {
	var parts []string
	for i := 0; i < len(id); i += n {
		var buffer bytes.Buffer
		for j := i; j < i+n; j++ {
			if j < len(id) {
				buffer.WriteString(string(id[j]))
			}
		}
		parts = append(parts, buffer.String())
	}
	return parts
}
