package main

import (
	"fmt"

	"github.com/kengru/problems/leetcode"
)

func main() {
	for _, v := range leetcode.Examples876.Tests {
		fmt.Println(leetcode.MiddleNode876(v))
	}
}
