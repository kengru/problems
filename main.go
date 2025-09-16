package main

import (
	"fmt"

	"github.com/kengru/problems/leetcode"
)

func main() {
	for _, v := range leetcode.Examples1089.Tests {
		leetcode.DuplicateZeros1089(v)
		fmt.Println(v)
	}
}
