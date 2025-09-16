package leetcode

// Given a binary array nums, return the maximum number of consecutive 1's in the array.

func FindMaxConsecutiveOnes485(nums []int) int {
	count := 0
	max := 0
	for _, n := range nums {
		if n == 1 {
			count++
		} else {
			count = 0
		}
		if count > max {
			max = count
		}
	}
	return max
}
