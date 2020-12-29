package Week_09

import "math"

//climb stairs
//1 || 2 ,different way
func climbStairs(n int )int {
	dp:=make([]int,n+1)
	dp[1]=1
	dp[2]=2
	for i:=3;i<=n;i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	return dp[n]
}
func climbStairs2(n int )int {
	if n<=2{
		return n
	}
	f1:=1
	f2:=2
	for i:=3;i<=n;i++{
		f1,f2=f2,f1+f2
	}
	return f2
}

// robs
func rob(nums []int) int {
	if len(nums)<1{
		return 0
	}
	dp:=make([][2]int,len(nums))
	dp[0][0]=0
	dp[0][1]=nums[0]
	for i:=1;i<len(nums);i++{
		dp[i][0]=maxRob(dp[i-1][1],dp[i-1][0])
		dp[i][1]=dp[i-1][0]+nums[i]
	}
	return maxRob(dp[len(nums)-1][0],dp[len(nums)-1][1])
}
func maxRob(a,b int )int {
	if a>b{
		return a
	}
	return b
}
func RobRing(nums []int )int {
	a:=rob(nums[1:])
	b:=rob(nums[:len(nums)-1])
	return maxRob(a,b)
}

func coinChange(coins []int,amount int)int{
	if len(coins)<1{
		return 0
	}
	dp:=make([]int,amount+1)
	dp[0]=0
	for i:=1;i<=amount;i++{
		dp[i]=-1
		for c:=0;c<len(coins);c++{
			if i<coins[c]||dp[i-coins[c]]==-1{
				continue
			}
			count:=dp[i-coins[c]]+1
			if dp[i]==-1||dp[i]>count{
				dp[i]=count
			}
		}
	}
	return dp[amount]
}

// 最大连续子数组的和
func maxSumArray(nums []int)int {
	if len(nums)<1{
		return 0
	}
	var max = nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i]<nums[i]+nums[i-1]{
			nums[i]=nums[i]+nums[i-1]
		}
		if nums[i]>max{
			max=nums[i]
		}
	}
	return max
}

// 跳跃游戏 45
func canJump(nums []int)int{
	dp:=make([]int,len(nums)+1)
	dp[0]=0
	for i:=1;i<len(nums);i++{
		dp[i]=math.MaxInt32
		for j:=0;j<i;j++ {
			if j+nums[j]>=i{
				dp[i] = minStep(dp[i],dp[j]+1)
			}
		}
	}
	return dp[len(dp)-1]
}

func minStep(a,b int )int {
	if a>b{
		return b
	}
	return a
}
// 不同路径
func pathUnique(m,n int )int{
	dp:=make([][]int,m)
	for i:=range dp{
		dp[i]=make([]int,n)
	}
	dp[0][0]=0
	for i:=1;i<m;i++{
		dp[i][0]=1
	}
	for j:=1;j<n;j++{
		dp[0][j]=1
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j]=dp[i-1][j]+dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
// 三角形最小路径和
//bottom to top //it's triangle,  so isn't need to consider range
func trianglePath(triangle [][]int)int {
	for i:=len(triangle)-2;i>=0;i--{
		for j:=0;j< len(triangle[i]);j++{
			triangle[i][j]=minNumber(triangle[i+1][j],triangle[i+1][j+1])+triangle[i][j]
		}
	}
	return triangle[0][0]
}
func minNumber(a,b int)int{
	if a>b{
		return b
	}
	return a
}


// 最长公共子序列 1143
func longestCommonSubsequence(ta,tb string )int {
	dp:=make([][]int,len(ta)+1)
	for i:=range dp{
		dp[i]=make([]int,len(tb)+1)
	}
	for i:=1;i<len(ta);i++{
		for j:=1;j<len(tb);j++{
			if ta[i-1] == tb[j-1]{
				dp[i][j]=dp[i-1][j-1]+1
			}else{
				dp[i][j]=maxSeq(dp[i-1][j],dp[i][j-1])+1
			}
		}
	}
	return dp[len(ta)][len(tb)]
}
func maxSeq(a, b int )int{
	if a>b{
		return a
	}
	return b
}

// 股票
func maxProfits(prices []int)int{
	sum:=0
	for i:=1;i<len(prices);i++{
		if prices[i]-prices[i-1]>0{
			sum+=prices[i]-prices[0]
		}
	}
	return sum
}
// 摆动序列