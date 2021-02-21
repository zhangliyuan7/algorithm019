package Repeat_01

import "math"

// 312 戳气球
func maxCoins(nums []int) int {
	ln:=len(nums)
	points:=make([]int,ln+2)
	points[0],points[ln+1]=1,1
	for i:=range nums{
		points[i+1]=nums[i]
	}
	//为了使其开区间，增加了两个边，令其为1 ，所以分数不会受影响
	dp:=make([][]int,ln+2)
	for i:=range dp{
		dp[i]=make([]int,ln+2)
	}
	// i,j 开区间
	// 戳破i-j 中的k 那么先要戳破 i-k 和k-j ，这样 等到戳破k的时候 左边就只有i，右边就只有j
	// dp[i][j]=dp[i][k]+dp[k][j]+points[i]*points[k]*points[j]
	// base case if i==j { dp[i][j]=0 }
	// 因为在这种情况下，开区间i j 中间没有气球可以戳破
	for i:=ln;i>=0;i-- {
		for j:=i+1;j<ln+2;j++{
			//递推方向是东北方向 ，所以行从大到小，列 从小到大，倒三角迭代
			for k:=i+1;k<j;k++{
				dp[i][j]=int(math.Max(
					float64(dp[i][k]+dp[k][j]+points[i]*points[k]*points[j]),
					float64(dp[i][j])),
				)
			}
		}
	}
	return dp[0][ln+1]
}