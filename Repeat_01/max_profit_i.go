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

//188
func maxProfitIV(k int, prices []int) int {
	ln:=len(prices)
	// 大于一半 即不限制次数
	if k>ln/2{
		return maxProfitInf(prices)
	}
	dp:=make([][][2]int,ln)
	for i:=range dp{
		dp[i]=make([][2]int,k+1)
	}
	for i:=0;i<k+1;i++{
		dp[0][i][1]=-prices[0]
		dp[0][i][0]=0
	}
	for i:=1;i<ln;i++{
		// 因为j从1开始(为了递推公式)，所以k次为<=k
		for j:=1;j<k+1;j++{
			// 手中无股票 可以来自第k次买入，抛售掉，或者之前就抛售掉了或者从未持有
			dp[i][j][0]=max(dp[i-1][j][1]+prices[i],dp[i-1][j][0])
			// 手中有股票，则是k-1次的手中没有股票的情况，或者之前一直持有，当前次数第k次
			// 如果当前持有股票来自于没有股票的情况，那么一定是[i-1][k-1][0]的情况
			dp[i][j][1]=max(dp[i-1][j-1][0]-prices[i],dp[i-1][j][1])
		}
	}
	return dp[ln-1][k][0]
}


func maxProfitInf(prices []int) int {
	if len(prices)<2{
		return 0
	}
	sum:=0
	plen:=len(prices)
	for i:=1;i<plen;i++{
		if prices[i]-prices[i-1]>0{
			sum+=prices[i]-prices[i-1]
		}
	}
	return sum
}

//309
func maxProfitsV(prices []int)int{
	ln:=len(prices)
	dp:=make([][2]int,ln)
	if ln<=1{
		return 0
	}
	if ln==2{
		if prices[1]-prices[0]>0{
			return prices[1]-prices[0]
		}
		return 0
	}

	dp[0][0]=0
	dp[0][1]=-prices[0]
	dp[1][0]=max(dp[0][0],prices[1]-prices[0])
	dp[1][1]=max(dp[0][1],-prices[1])
	dp[2][0]=max(dp[1][0],dp[0][1]+prices[2])
	dp[2][1]=max(dp[1][1],dp[0][0]-prices[2])
	for i:=2;i<ln;i++{
		dp[i][0]=max(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1]=max(dp[i-1][1],dp[i-2][0]-prices[i])
	}
	return dp[ln-1][0]
}

// 714
func maxProfitVI(prices []int, fee int) int {
	ln:=len(prices)
	dp:=make([][2]int,ln)
	dp[0][0]=0
	dp[0][1]=-prices[0]
	for i:=1;i<ln;i++{
		dp[i][0]=max(dp[i-1][0],dp[i-1][1]-fee+prices[i])
		dp[i][1]=max(dp[i-1][1],dp[i][0]-prices[i])
	}
	return dp[ln-1][0]
}
