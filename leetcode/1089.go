package leetcode

// Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

// Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.

var Examples1089 = Example[[]int]{
	Tests: [][]int{
		{1, 0, 2, 3, 0, 4, 5, 0},
		{1, 0, 0, 2, 3, 0, 0, 4},
	},
}

func DuplicateZeros1089(arr []int) {
	idx1 := 0
	for idx1 < len(arr)-1 {
		if arr[idx1] == 0 {
			shiftRight(arr, idx1)
			idx1 += 2
		} else {
			idx1++
		}
	}
}

func shiftRight(arr []int, idx int) {
	for i := len(arr) - 2; i >= idx; i-- {
		arr[i+1] = arr[i]
	}
}
