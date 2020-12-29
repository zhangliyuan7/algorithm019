package Week_09

import "math"

// 188
func maxProfit(k int, prices []int) int {
	if k>len(prices)/2{
		return maxProfitInf(prices)
	}
	dp_i_0, dp_i_1 := make([]int, k+1), make([]int, k+1)
	for i := 0; i < k+1; i++{
		dp_i_0[i] = 0
		dp_i_1[i] = math.MinInt64
	}
	for _, v := range prices{
		for i := k; i >=1; i--{
			dp_i_0[i] = max(dp_i_0[i], dp_i_1[i] + v)
			dp_i_1[i] = max(dp_i_1[i], dp_i_0[i-1] - v)
		}
	}
	return dp_i_0[k]
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
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
// 不限制交易次数的最大利润dp
func MaxProfit(prices []int)int{
//	0 sell 1 buy
	dp:=make([][2]int,len(prices))
	dp[0][0]=0
	dp[0][1]=-prices[0]
	for i:=1;i<len(prices);i++{
		dp[i][0]=max(prices[i]-dp[i-1][1],dp[i-1][0])
		dp[i][1]=dp[i-1][0]-prices[i]
	}
	return dp[len(prices)-1][0]
}
//
func MaxProfitOnce(prices []int)int{
	minp:=math.MaxInt32
	minprice:=minp
	maxprofit:=0
	for i:=0;i<len(prices);i++{
		maxprofit=max(prices[i]-minprice,maxprofit)
		minprice=min(prices[i],minprice)
	}
	return maxprofit
}
