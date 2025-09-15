package leetcode

// Given an integer num, return the number of steps to reduce it to zero.

// In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.

var Examples1342 = Example[int]{
	Tests: []int{
		14,
		8,
		123,
	},
}

func NumberOfSteps1342(num int) int {
	steps := 0
	for num > 0 {
		if num&1 == 0 {
			num >>= 1
		} else {
			num -= 1
		}
		steps++
	}
	return steps
}
