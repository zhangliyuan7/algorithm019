package Repeat_01

import "sort"

// 628
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	length:=len(nums)
	if nums[0]>=0|| nums[length-1]<=0{
		return nums[length-1]*nums[length-2]*nums[length-3]
	}
	cand1:=nums[length-1]*nums[length-2]*nums[length-3]
	cand2:=nums[0]*nums[1]*nums[length-1]
	if cand1 >cand2{
		return cand1
	}
	return cand2
}
