package Repeat_01

import "math"

// 121 一次交易
func maxProfit(prices []int)int{
	maxP:=0
	minP:=math.MaxInt32
	// 暴力法，超时
	// for i:=0;i<len(prices);i++{
	// 	for j:=i+1;j<len(prices);j++{
	// 		if prices[j]-prices[i]>mp{
	// 			mp=prices[j]-prices[i]
	// 		}
	// 	}
	// }
	for i:=0;i<len(prices);i++{
		maxP=max(maxP,prices[i]-minP)
		minP=min(minP,prices[i])
	}
	return maxP
}
// 两个状态分别为 手上有股票，手上没有股票的当前最大利润
func maxProfitDP(prices []int)int{
	ln:=len(prices)
	dp:=make([][2]int,ln)
	dp[0][1]=-prices[0]
	dp[0][0]=0
	for i:=1;i<ln;i++{
		dp[i][0]=max(dp[i-1][1]+prices[i],dp[i-1][0])
		dp[i][1]=max(dp[i-1][1],-prices[i])
	}
	return dp[ln-1][0]
}

// 122 不限制次数
func maxProfitII(prices []int)int{
	mp:=0
	for i:=1;i<len(prices);i++{
		if prices[i]-prices[i-1]>0{
			mp+=prices[i]-prices[i-1]
		}
	}
	return mp
}
// 123
func maxProfitIII(prices []int)int {
		ln:=len(prices)
		dp:=make([][][2]int,ln)
		for i:=0;i<ln;i++{
		dp[i]=make([][2]int,3)
	}
	// 初始化
	// 因为不限制当天操作次数，所以当天可以完成完整两次操作
	// 所以要初始化所有操作次数的第一天情形
	dp[0][1][0], dp[0][1][1] = 0, -prices[0]
	dp[0][2][0], dp[0][2][1] = 0, -prices[0]
	for i:=1;i<ln;i++{
		// 第一次操作持有(buy)  第一次操作第一步
		dp[i][1][1]=max(dp[i-1][1][1],-prices[i])
		// 第一次操作且不持有(sell)  第一次操作第二步
		dp[i][1][0]=max(dp[i-1][1][1]+prices[i],dp[i-1][1][0])
		// 第二次操作 持有（buy） 第二次操作第一步
		dp[i][2][1]=max(dp[i-1][1][0]-prices[i],dp[i-1][2][1])
		// 第二次操作 不持有 （sell） 即完成完整第二次操作
		dp[i][2][0]=max(dp[i-1][2][0],dp[i-1][2][1]+prices[i])
	}
	return dp[ln-1][2][0]
}

//