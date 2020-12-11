package Week_06

// 198
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//define dp list
	dp:=make([][2]int,len(nums))
	dp[0][0]=0
	dp[0][1]=nums[0]
	for i:=1;i<len(nums);i++ {
		//可以连续两个都不抢
		dp[i][0]=max_rob(dp[i-1][1],dp[i-1][0])
		dp[i][1]=dp[i-1][0]+nums[i]
	}

	return max_rob(dp[len(nums)-1][0],dp[len(nums)-1][1])
}

//化简子问题，nums[i]=i-1/i-2+nums[i] 因为i-2包含了i-3，i-4等等条件，所以递推公式成立
func robA(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp :=make([]int, len(nums))
	dp[0]=nums[0]
	dp[1]=max_rob(nums[0],nums[1])
	for i:=2;i<len(nums);i++{
		dp[i]=max_rob(dp[i-1],dp[i-2]+nums[i])
	}
	return dp[len(dp)-1]
}
// robII 213
func Rob213(nums []int)int {
	//nums[0]与末尾只有一个可以抢或者都不抢
	//拆分成从0到尾部-1或从1到尾部两个dp
	if len(nums)==0{
		return 0
	}
	if len(nums)==1{
		return nums[0]
	}
	if len(nums)==2{
		return max_rob(nums[0],nums[1])
	}
	dp1:=make([]int,len(nums))
	dp2:=make([]int,len(nums))
	dp1[0]=nums[0]
	dp1[1]=max_rob(nums[0],nums[1])
	for i:=2;i<len(nums)-1;i++{
		dp1[i]=max_rob(dp1[i-1],dp1[i-2]+nums[i])
	}
	dp2[1]=nums[1]
	dp2[2]=max_rob(nums[1],nums[2])
	for i:=3;i<len(nums);i++{
		dp2[i]=max_rob(dp2[i-1],dp2[i-2]+nums[i])
	}
	return max_rob(dp2[len(dp2)-1],dp1[len(dp1)-2])
}
func max_rob(a,b int) int {
	if a < b {
		return b
	}
	return a
}