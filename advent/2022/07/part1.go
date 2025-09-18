package advent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kengru/problems/advent"
)

func Year20220701() {
	// Getting input
	lines := advent.GetLinesFromFile("advent/2022/07/ie.txt")
	// lines := advent.GetLinesFromFile("advent/2022/07/input.txt")

	dirs := []*advent.Directory{}
	stack := advent.Stack[*advent.Directory]{}
	current := &advent.Directory{}
	idx := 0
	for idx < len(lines) {
		line := lines[idx]
		if line[0] == '$' {
			cmd := strings.Split(line, " ")
			if cmd[1] == "cd" {
				if cmd[2] != ".." {
					dir := advent.Directory{
						Name:  cmd[2],
						Size:  0,
						Dirs:  []*advent.Directory{},
						Files: []*advent.File{},
					}
					dirs = append(dirs, &dir)
					current = &dir
					stack.Add(&dir)
				} else {
					current = stack.Pop()
				}
			}
			idx++
			continue
		}
		info := strings.Split(line, " ")
		if info[0] == "dir" {
			dir := advent.Directory{
				Name:  info[1],
				Size:  0,
				Dirs:  []*advent.Directory{},
				Files: []*advent.File{},
			}
			current.Dirs = append(current.Dirs, &dir)
		} else {
			size, _ := strconv.Atoi(info[0])
			file := advent.File{
				Name: info[1],
				Size: size,
			}
			current.Size += size
			current.Files = append(current.Files, &file)
		}
		fmt.Println(lines[idx])
		idx++
	}

	fmt.Println(dirs)
	fmt.Println(current.Name)
}
