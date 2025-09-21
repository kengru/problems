package leetcode

// Given a binary array nums, return the maximum number of consecutive 1's in the array if you can flip at most one 0.

var Examples487 = Example[[]int]{
	Tests: [][]int{
		{1, 0, 1, 1, 0},
		{1, 0, 1, 1, 0, 1},
		{1, 1, 0, 1},
	},
}

func FindMaxConsecutiveOnesII487(nums []int) int {
	idx1, idx2 := 0, 0
	max := 0

	for idx1 < len(nums) {
		zeroes := 0
		ones := 0
		for idx2 < len(nums) && zeroes < 2 {
			value := nums[idx2]
			if value == 0 {
				zeroes++
			}
			ones++
			idx2++
		}
		if idx2 == len(nums) && zeroes < 2 {
			if ones > max {
				max = ones
			}
		} else {
			if ones-1 > max {
				max = ones - 1
			}
		}
		idx1++
		idx2 = idx1
	}

	return max
}
