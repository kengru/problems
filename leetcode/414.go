package leetcode

import (
	"slices"
)

// Given an integer array nums, return the third distinct maximum number in this array.
// If the third maximum does not exist, return the maximum number.

var Examples414 = Example[[]int]{
	Tests: [][]int{
		{3, 2, 1},
		{1, 2},
		{2, 2, 3, 1},
	},
}

func ThirdMax414(nums []int) int {
	slices.Sort(nums)
	biggest := nums[len(nums)-1]
	times := 1
	i := len(nums) - 2
	for i >= 0 && times != 3 {
		if nums[i] < biggest {
			times++
			biggest = nums[i]
		}
		i--
	}
	if times < 3 {
		return nums[len(nums)-1]
	}
	return biggest
}
