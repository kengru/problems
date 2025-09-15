package leetcode

// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

// Each letter in magazine can only be used once in ransomNote.

func CanConstruct383(ransomNote string, magazine string) bool {
	mag := make(map[rune]int)
	for _, v := range magazine {
		_, ex := mag[v]
		if !ex {
			mag[v] = 1
		} else {
			mag[v]++
		}
	}
	for _, v := range ransomNote {
		_, ex := mag[v]
		if !ex {
			return false
		}
		if mag[v] == 0 {
			return false
		}
		mag[v]--
	}
	return true
}
