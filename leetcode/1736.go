package leetcode

import (
	"strconv"
	"strings"
)

// Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the remaining elements to the right.

// Note that elements beyond the length of the original array are not written. Do the above modifications to the input array in place and do not return anything.

var Examples1736 = Example[string]{
	Tests: []string{
		"2?:?0",
		"0?:3?",
		"1?:22",
	},
}

func MaximumTime1736(time string) string {
	sp := strings.Split(time, "")

	if sp[0] == "?" {
		if sp[1] != "?" {
			val, _ := strconv.Atoi(sp[1])
			if val > 3 {
				sp[0] = "1"
			} else {
				sp[0] = "2"
			}
		} else {
			sp[0] = "2"
		}
	}

	if sp[1] == "?" {
		if sp[0] != "2" {
			sp[1] = "9"
		} else {
			sp[1] = "3"
		}
	}

	if string(sp[3]) == "?" {
		sp[3] = "5"
	}
	if string(sp[4]) == "?" {
		sp[4] = "9"
	}
	return strings.Join(sp, "")
}
