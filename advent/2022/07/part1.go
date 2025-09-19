package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220701() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2022/07/ie.txt")
	lines := advent.GetLinesFromFile("advent/2022/07/input.txt")

	dirs := []*advent.Directory{}
	current := &advent.Directory{}
	idx := 0
	for idx < len(lines) {
		line := lines[idx]
		if line[0] == '$' {
			cmd := strings.Split(line, " ")
			if cmd[1] == "cd" {
				if cmd[2] != ".." {
					dir := advent.Directory{
						Parent: current,
						Name:   cmd[2],
						Size:   0,
					}
					dirs = append(dirs, &dir)
					current = &dir
				} else {
					current = current.Parent
				}
			}
			idx++
			continue
		}
		info := strings.Split(line, " ")
		if info[0] != "dir" {
			size, _ := strconv.Atoi(info[0])
			current.Size += size
			parent := current.Parent
			for parent != nil {
				parent.Size += size
				parent = parent.Parent
			}
		}
		idx++
	}

	total := 0
	for _, d := range dirs {
		if d.Size <= 100000 {
			total += d.Size
		}
	}

	fmt.Println(total)
}
