package Repeat_01

import "math"

// 53

func maxSubArray(nums []int) int {
	for i:=1;i<len(nums);i++{
		if nums[i]<nums[i-1]+nums[i]{
			nums[i]+=nums[i-1]
		}
	}
	max:=math.MinInt32
	for i:=range nums{
		if nums[i]>max{
			max=nums[i]
		}
	}
	return max
}
