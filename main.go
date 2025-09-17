package main

import (
	"fmt"

	"github.com/kengru/problems/leetcode"
)

func main() {
	for _, v := range leetcode.Examples448.Tests {
		// leetcode.SortArrayByParity905(v)
		fmt.Println(leetcode.FindDisappearedNumbers448(v))
	}
	// [0,1,4,0,3]
}
