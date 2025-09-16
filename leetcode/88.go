package leetcode

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
// representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
// To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

func Merge88(nums1 []int, m int, nums2 []int, n int) {
	idx1 := m - 1
	idx2 := n - 1
	result_idx := m + n - 1
	for idx2 >= 0 {
		if idx1 >= 0 && nums1[idx1] > nums2[idx2] {
			nums1[result_idx] = nums1[idx1]
			idx1 -= 1
		} else {
			nums1[result_idx] = nums2[idx2]
			idx2 -= 1
		}
		result_idx--
	}
}
