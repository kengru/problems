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
	i, j := 0, 1
	for i < len(arr)-1 {
		greatest := arr[j]
		for j < len(arr) {
			if arr[j] > greatest {
				greatest = arr[j]
			}
			j++
		}
		arr[i] = greatest
		i++
		j = i + 1
	}
	arr[len(arr)-1] = -1
	return arr
}
