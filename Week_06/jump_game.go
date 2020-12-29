package Week_06

import (
	"math"
)

/// hen man 45
func Jump(nums []int) int {
	var dp = make([]int,len(nums))
	dp[0]=0
	for i:=1;i<len(nums);i++{
		dp[i]=math.MaxInt64
		for j:=0;j<i;j++{
			if nums[j]+j>=i{
				dp[i]=min(dp[i],dp[j]+1)
			}
		}
	}
	return dp[len(dp)-1]
}
func min(a,b int )int{
	if a>b{
		return b
	}
	return a
}