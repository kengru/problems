package main

import (
	"fmt"

	"github.com/kengru/problems/leetcode"
)

func main() {
	for _, v := range leetcode.Examples1299.Tests {
		fmt.Println(leetcode.ReplaceElements1299(v))
		// fmt.Println(leetcode.MaximumTime1736(v))
	}
	// [0,1,4,0,3]
}
