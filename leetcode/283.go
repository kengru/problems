package leetcode

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

var Examples283 = Example[[]int]{
	Tests: [][]int{
		{0, 1, 0, 3, 12},
		{0},
	},
}

func MoveZeroes283(nums []int) {
	if len(nums) < 2 {
		return
	}
	i := 0
	j := 1
	for i < len(nums) {
		if nums[i] == 0 {
			for j < len(nums) {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
				j++
			}
		}
		i++
		j = i + 1
	}
}
