package leetcode

// Given an array arr of integers, check if there exist two indices i and j such that :

// i != j
// 0 <= i, j < arr.length
// arr[i] == 2 * arr[j]

var Examples1346 = Example[[]int]{
	Tests: [][]int{
		{10, 2, 5, 3},
		{3, 1, 7, 11},
		{0, -2, 2},
	},
}

func CheckIfExist1346(arr []int) bool {
	if len(arr) < 2 {
		return false
	}
	mapa := make(map[int]int)
	for i := range len(arr) {
		mapa[arr[i]] = i
	}

	idx := 0
	for idx < len(arr) {
		double := arr[idx] * 2
		j, ex := mapa[double]
		if ex && idx != j {
			return true
		}
		if arr[idx]&1 == 0 {
			j, ex = mapa[arr[idx]/2]
			if ex && idx != j {
				return true
			}
		}
		idx++
	}

	return false
}
