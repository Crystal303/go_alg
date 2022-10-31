package dp

// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999

// length 代表台阶的高度 [0,n]

// minCostClimbingStairs 最小代价爬楼梯
func minCostClimbingStairs(cost []int) int {
	c0 := cost[0]
	c1 := cost[1]
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	l := len(cost)
	if l == 2 {
		return min(c0, c1)
	}
	if l == 3 {
		return min(min(c0, c1)+cost[2], c1)
	}

	c0 = min(cost[0], cost[1])    // 2
	c1 = min(cost[1], c0+cost[2]) // 3
	for i := 4; i <= l; i++ {
		temp := c1
		c1 = min(c0+cost[i-2], c1+cost[i-1])
		c0 = temp
	}
	return c1
}

// fib 斐波那契数列
func fib(n int) int {
	if n == 0 ||
		n == 1 {
		return n
	}
	f0 := 0
	f1 := 1
	num := 0
	for i := 2; i <= n; i++ {
		num = f0 + f1
		f0 = f1
		f1 = num
	}
	return num
}

// isSubsequence 子序列
func isSubsequence(s string, t string) bool {
	sl, tl := len(s), len(t)
	if tl < sl {
		return false
	}
	if tl == sl {
		return s == t
	}
	if sl == 0 {
		return true
	}

	si := 0
	for i := 0; i < tl; i++ {
		if t[i] == s[si] {
			si++
		}
		if sl <= si {
			return true
		}
	}
	return false
}

// maxProfit 股票最大价格
// [7,1,5,3,6,4]
func maxProfit(prices []int) int {
	min, max := prices[0], prices[0]
	l := len(prices)
	if l == 1 {
		return 0
	}

	wave := 0
	for i := 1; i < l; i++ {
		if p := prices[i]; p < min {
			if w := max - min; wave < w {
				wave = w
			}
			min = p
			max = p
		} else {
			if min < p && max < p {
				max = p
			}
		}
	}
	if w := max - min; wave < w {
		wave = w
	}

	return wave
}

// climbStairs 爬楼梯
// 1 <= n <= 45
// Yn := Yn-1  + Yn-2
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	f := 0
	f1, f2 := 1, 2
	for i := 3; i <= n; i++ {
		f = f1 + f2
		f1 = f2
		f2 = f
	}
	return f
}

// generate 杨辉三角生成
func generate(numRows int) [][]int {
	// 初始化
	ret := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		l := i + 1
		item := make([]int, l)
		item[0] = 1
		item[i] = 1
		ret[i] = item

		if 2 < l {
			for j := 1; j < i; j++ {
				item[j] = ret[i-1][j] + ret[i-1][j-1]
			}
		}
	}

	return ret
}

// generate 杨辉三角生成 单独的第几行
func getRow(rowIndex int) []int {
	ret := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		ret[0] = 1
		ret[i] = 1

		// 1 2 1
		// 1 3 3 1
		if l := i + 1; 2 < l {
			temp := make([]int, i)
			copy(temp, ret)
			for j := 1; j < i; j++ {
				ret[j] = temp[j] + temp[j-1]
			}
		}
	}
	return ret
}
