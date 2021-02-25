package Repeat_01

// 152
//maxDP[i + 1] = max(maxDP[i] * A[i + 1], A[i + 1],minDP[i] * A[i + 1])
//minDP[i + 1] = min(minDP[i] * A[i + 1], A[i + 1],maxDP[i] * A[i + 1])
//
func maxProductB(nums []int) int {
	ln:=len(nums)
	max:=nums[0]
	maxdp:=make([]int,ln)
	mindp:=make([]int,ln)
	maxdp[0]=nums[0]
	mindp[0]=nums[0]
	for i:=1;i<ln;i++{
		maxdp[i]=maxp(maxdp[i-1]*nums[i],maxp(nums[i],mindp[i-1]*nums[i]))
		mindp[i]=minp(mindp[i-1]*nums[i],minp(nums[i],maxdp[i-1]*nums[i]))
		max=maxp(maxdp[i],max)
	}
	return max
}

func maxProductA(nums []int) int {
	ans:=nums[0]
	maxF:=nums[0]
	minF:=nums[0]
	for i:=1;i<len(nums);i++{
		mx,mn:=maxF,minF
		// 当前乘到的位置的最大值
		maxF=maxp(mx*nums[i],maxp(nums[i],nums[i]*mn))
		// 最小值，当前乘到的位置的最小值， 遇到负数，最小值有可能变为最大值
		// 当遇到0的时候
		minF=minp(mn*nums[i],minp(nums[i],nums[i]*mx))
		ans=maxp(maxF,ans)
	}
	return ans
}
func maxp(a,b int )int {
	if a>b{
		return a
	}
	return b
}
func minp(a,b int )int {
	if a<b{
		return a
	}
	return b
}