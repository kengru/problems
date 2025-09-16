package leetcode

type Example[K any] struct {
	Tests []K
}

func shiftLeft(nums []int, idx int) {
	for i := idx; i < len(nums); i++ {
		if i+1 == len(nums) {
			nums[i] = -9999
		} else {
			nums[i] = nums[i+1]
		}
	}
}
