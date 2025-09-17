package leetcode

import (
	"slices"
)

// A school is trying to take an annual photo of all the students.
// The students are asked to stand in a single file line in non-decreasing order by height.
// Let this ordering be represented by the integer array expected where expected[i] is the expected height of the ith student in line.

// You are given an integer array heights representing the current order that the students are standing in.
// Each heights[i] is the height of the ith student in line (0-indexed).

// Return the number of indices where heights[i] != expected[i].

var Examples1051 = Example[[]int]{
	Tests: [][]int{
		{1, 1, 4, 2, 1, 3},
		{5, 1, 2, 3, 4},
		{1, 2, 3, 4, 5},
	},
}

func HeightChecker1051(heights []int) int {
	diff := 0
	fixed := slices.Clone(heights)
	slices.Sort(fixed)
	for i := range heights {
		if heights[i] != fixed[i] {
			diff++
		}
	}
	return diff
}
