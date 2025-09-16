package leetcode

// Given a binary string s ​​​​​without leading zeros, return true​​​ if s contains at most one contiguous segment of ones. Otherwise, return false.

func CheckOnesSegment1784(s string) bool {
	segments := 0
	onSegment := false
	for _, v := range s {
		if string(v) == "1" {
			if !onSegment {
				segments++
			}
			onSegment = true
		} else {
			onSegment = false
		}
	}
	return segments == 1
}
