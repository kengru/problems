package leetcode

// Given an array of integers arr, return true if and only if it is a valid mountain array.

// Recall that arr is a mountain array if and only if:

// arr.length >= 3
// There exists some i with 0 < i < arr.length - 1 such that:
// arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
// arr[i] > arr[i + 1] > ... > arr[arr.length - 1]

var Examples941 = Example[[]int]{
	Tests: [][]int{
		{2, 1},
		{3, 5, 5},
		{0, 3, 2, 1},
		{2, 0, 2},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{0, 1, 2, 1, 2},
	},
}

func ValidMountainArray941(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	rising := false
	changes := 1
	if arr[0] < arr[1] {
		rising = true
	}
	i := 0
	for i < len(arr)-1 {
		if arr[i] == arr[i+1] {
			return false
		}
		if rising && changes&1 != 0 && arr[i] > arr[i+1] {
			changes++
			rising = false
		}
		if !rising && changes&1 == 0 && arr[i] < arr[i+1] {
			changes++
			rising = true
		}
		i++
	}

	return changes == 2
}
