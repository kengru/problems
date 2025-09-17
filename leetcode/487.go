package leetcode

// Given an integer array nums, return the third distinct maximum number in this array.
// If the third maximum does not exist, return the maximum number.

var Examples487 = Example[[]int]{
	Tests: [][]int{
		{4, 3, 2, 7, 8, 2, 3, 1},
		{1, 1},
	},
}

func FindMaxConsecutiveOnes448(nums []int) []int {
	full := make(map[int]bool)
	for i := range len(nums) {
		_, ex := full[i+1]
		if !ex {
			full[i+1] = true
		}
	}
	for _, v := range nums {
		delete(full, v)
	}
	keys := make([]int, 0, len(full))
	for k := range full {
		keys = append(keys, k)
	}
	return keys
}
