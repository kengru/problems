package leetcode

import "strconv"

// Given an integer n, return a string array answer (1-indexed) where:

// answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
// answer[i] == "Fizz" if i is divisible by 3.
// answer[i] == "Buzz" if i is divisible by 5.
// answer[i] == i (as a string) if none of the above conditions are true.

var Examples420 = Example[int]{
	Tests: []int{
		3,
		5,
		15,
	},
}

func FizzBuzz420(n int) []string {
	result := []string{}
	for i := 1; i <= n; i++ {
		fb := ""
		if i%3 == 0 {
			fb += "Fizz"
		}
		if i%5 == 0 {
			fb += "Buzz"
		}
		if fb == "" {
			fb = strconv.Itoa(i)
		}
		result = append(result, fb)
	}
	return result
}
