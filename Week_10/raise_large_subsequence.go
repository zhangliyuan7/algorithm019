package Week_10

// 300
func lengthOfLIS(nums []int) int {
	dp :=make([]int,len(nums))
	for i:=range dp{
		dp[i]=1
	}
	for i:=1;i<len(nums);i++{
		for j:=0;j<i;j++{
			if nums[j]<nums[i] {
				dp[i] = maxLength(dp[i], dp[j]+1)
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
func maxLength(a,b int )int{
	if a>b{
		return a
	}
	return b
}