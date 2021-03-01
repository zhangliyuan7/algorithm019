package Repeat_01

// 303
type NumArray struct {
	dp []int
	nums []int
	ln int
}


func ConstructorSA(nums []int) NumArray {
	ln:=len(nums)
	if ln==0{
		return NumArray{}
	}
	dp:=make([]int,ln)
	dp[0]=nums[0]
	for i:=1;i<len(nums);i++{
		dp[i]=dp[i-1]+nums[i]
	}
	na:=NumArray{
		dp:dp,
		nums:nums,
		ln:ln,
	}
	return na
}


func (this *NumArray) SumRange(i int, j int) int {
	if this.ln==0{
		return 0
	}
	if j>=this.ln{
		j=this.ln-1
	}
	if i<0{
		i=0
	}
	return this.dp[j]-this.dp[i]+this.nums[i]
}


