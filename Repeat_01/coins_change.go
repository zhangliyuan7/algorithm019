package Repeat_01

//322
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for am := 0; am <= amount; am++ {
		for c := 0; c < len(coins); c++ {
			if coins[c] > am {
				continue
			}
			dp[am] = min(dp[am], dp[am-coins[c]]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func coinChangeI(coins []int, amount int) int {
	if len(coins) < 1 || amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for c := 0; c < len(coins); c++ {
			if i < coins[c] || dp[i-coins[c]] == -1 {
				continue
			}
			count := dp[i-coins[c]] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	return dp[amount]
}

//1 外层需要是coins循环，否则结果是排列数 而非组合数
//2 内层从当前的coin[i]值开始 ，否则会重复
//3 dp[i]=dp[i-5]+dp[i-2]+dp[i-1]
func change(amount int, coins []int) int {
	if amount == 0 || len(coins) < 0 {
		return 1
	}
	dp := make([]int, amount+1)
	dp[0] = 1
	// 必须硬币在外层循环 这样结果才是组合数
	// 如果不是这样 结果为排列数
	// dp[i]=dp[i-5]+dp[i-2]+dp[i-1]
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] += dp[i-c]
		}
	}
	return dp[amount]
}
