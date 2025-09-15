package leetcode

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
// Return the running sum of nums.

var Examples1480 = Example[[]int]{
	Tests: [][]int{
		{1, 2, 3, 4},
		{3, 1, 2, 10, 1},
	},
}

func RunningSum1480(nums []int) []int {
	s := 0
	for i, v := range nums {
		s += v
		nums[i] = s
	}
	return nums
}
