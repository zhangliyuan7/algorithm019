package Week_06

// 动态规划 只要往前累计比当前大，那么把之前的和加到当前来
func MaxSubArray(nums []int) int {
	var max int = nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i]+nums[i-1]>nums[i]{
			nums[i]=nums[i-1]+nums[i]
		}
		if nums[i]>max{
			max=nums[i]
		}
	}
	return max
}

func  MaxSubArray2(nums []int) int {
	m:=nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i]+nums[i-1]>nums[i]{
			nums[i]=nums[i]+nums[i-1]
		}
		if nums[i]>m{
			m=nums[i]
		}
	}
	return m
}

func MaxSubArrayT1(nums []int) int {
	m:=nums[0]
	for i:=1;i<len(nums);i++{
		if nums[i]+nums[i-1]>nums[i]{
			nums[i]=nums[i]+nums[i-1]
		}
		if nums[i]>m{
			m=nums[i]
		}
	}
	return m
}