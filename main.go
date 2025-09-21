package main

import (
	"fmt"

	"github.com/kengru/problems/leetcode"
)

// Leetcode Main
func main() {
	for _, v := range leetcode.Examples487.Tests {
		// leetcode.SortArrayByParity905(v)
		fmt.Println(leetcode.FindMaxConsecutiveOnesII487(v))
	}
}

// Advent Of Code Main
// func main() {
// 	advent.Year20220901()
// }
