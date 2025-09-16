package main

import (
	"fmt"

	"github.com/kengru/problems/leetcode"
)

func main() {
	for _, v := range leetcode.Examples1736.Tests {
		fmt.Println(leetcode.MaximumTime1736(v))
	}
}
