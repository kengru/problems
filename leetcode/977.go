package leetcode

// Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.

func SortedSquares977(nums []int) []int {
	result := make([]int, len(nums))
	idx1 := 0
	idx2 := len(nums) - 1
	res_idx := len(nums) - 1
	for idx2 >= idx1 {
		left := nums[idx1] * nums[idx1]
		right := nums[idx2] * nums[idx2]
		if right > left {
			result[res_idx] = right
			idx2 -= 1
		} else {
			result[res_idx] = left
			idx1 += 1
		}
		res_idx--
	}
	return result
}
