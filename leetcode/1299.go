package leetcode

// Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

// After doing so, return the array.

var Examples1299 = Example[[]int]{
	Tests: [][]int{
		{17, 18, 5, 4, 6, 1},
		{400},
	},
}

func ReplaceElements1299(arr []int) []int {
	max := -1
	i := len(arr) - 1
	for i >= 0 {
		if i+1 == len(arr) || arr[i] > max {
			temp := arr[i]
			arr[i] = max
			max = temp
			i--
			continue
		}
		arr[i] = max
		i--
	}
	return arr
}
