package advent

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220802() {
	// Getting input
	// lines := advent.GetLinesFromFile("advent/2022/08/ie.txt")
	lines := advent.GetLinesFromFile("advent/2022/08/input.txt")

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

	total := dirs[0].Size
	sizes := []int{}
	for _, d := range dirs {
		unused := 70000000 - total
		dif := unused + d.Size
		if dif >= 30000000 {
			sizes = append(sizes, d.Size)
		}
	}

	sort.Ints(sizes)
	fmt.Println(sizes[0])
}
