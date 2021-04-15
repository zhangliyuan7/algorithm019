package Repeat_01


//213 打家劫舍2
func robII(nums []int )int{
	ln:=len(nums)
	if ln==0{
		return 0
	}
	if ln==1{
		return nums[0]
	}
	// 不抢第一个 || 不抢最后一个 中最大的
	return max(robBase(nums[1:]),robBase(nums[:ln-1]))
}
// 打家劫舍1
func robBase(nums []int )int {
	ln:=len(nums)
	if ln<1{
		return 0
	}
	dp:=make([][2]int,ln)
	dp[0][0]=0
	dp[0][1]=nums[0]
	for i:=1;i<ln;i++{
		dp[i][0]=max(dp[i-1][0],dp[i-1][1])
		dp[i][1]=dp[i-1][0]+nums[i]
	}
	return max(dp[ln-1][0],dp[ln-1][1])
}