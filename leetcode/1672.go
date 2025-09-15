package leetcode

// You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.
// A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.

var Examples1672 = Example[[][]int]{
	Tests: [][][]int{
		{{1, 2, 3}, {3, 2, 1}},
		{{1, 5}, {7, 3}, {3, 5}},
	},
}

func MaximumWealth1672(accounts [][]int) int {
	maxWealth := 0
	for _, person := range accounts {
		total := 0
		for _, value := range person {
			total += value
		}
		if total > maxWealth {
			maxWealth = total
		}
	}
	return maxWealth
}
