package dp

// bag = 4
// w 2 3 4
// v 10 30 35

// 递推公式
// y = max{ f(y-i), f(y-Wi) + Vi}
//
//

// Bag01 0-1背包问题
func Bag01(weight []int, value []int, bag int) int {
	dp := make([][]int, len(weight)) // i 物品
	for k := range dp {
		dp[k] = make([]int, bag+1) // j 重量
	}
	// 初始化第一行
	for i := 0; i < 1; i++ {
		for j := 1; j <= bag; j++ {
			if weight[i] <= j {
				dp[i][j] = value[i]
			}
		}
	}

	// 遍历
	for i := 1; i < len(weight); i++ {
		for j := 1; j <= bag; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bag]
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
