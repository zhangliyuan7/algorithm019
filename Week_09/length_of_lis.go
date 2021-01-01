package Week_09

// 300
func lengthOfLIS(nums []int) int {
	dp:=make([]int,len(nums))
	for i:=0;i<len(dp);i++{
		dp[i]=1
	}
	for i:=0;i<len(nums);i++ {
		for j:=0;j<i;j++{
			if nums[j]<nums[i]{
				dp[i]=maxx(dp[i],dp[j]+1)
			}
		}
	}
	max:=0
	for i:=range dp{
		if dp[i]>max{
			max=dp[i]
		}
	}
	return max
}
func maxx(a,b int )int{
	if a>b{
		return a
	}
	return b
}