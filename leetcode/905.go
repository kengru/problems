package leetcode

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

var Examples905 = Example[[]int]{
	Tests: [][]int{
		{3, 1, 2, 4},
		{0},
	},
}

func SortArrayByParity905(nums []int) []int {
	even := []int{}
	odd := []int{}
	for i := range nums {
		if nums[i]&1 == 0 {
			even = append(even, nums[i])
		} else {
			odd = append(odd, nums[i])
		}
	}
	even = append(even, odd...)
	return even
}
