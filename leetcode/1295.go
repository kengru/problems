package leetcode

// Given an array nums of integers, return how many of them contain an even number of digits.

func FindNumbers1295(nums []int) int {
	even := 0
	for _, num := range nums {
		digits := 0
		for num > 0 {
			num /= 10
			digits++
		}
		if digits&1 == 0 {
			even++
		}
	}
	return even
}
