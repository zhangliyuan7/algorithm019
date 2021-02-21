package Repeat_01

import "math"

//1438
// 双单调队列，维护单调性，其他就是常规滑动窗口
func longestSubarrayRightWay(nums []int,limit int )int{
	maxQ:=[]int{}
	minQ:=[]int{}
	maxLength:=0
	l,r:=0,0
	for r<len(nums) {
		//维护单调递减队列，队首最大值
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < nums[r] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, nums[r])
		//维护单调递增队列 队首最小
		for len(minQ) > 0 && minQ[len(minQ)-1] > nums[r] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, nums[r])
		for len(maxQ) > 0 && len(minQ) > 0 && maxQ[0]-minQ[0] > limit{
			if nums[l]==maxQ[0]{
				maxQ=maxQ[1:]
			}
			if nums[l]==minQ[0]{
				minQ=minQ[1:]
			}
			l++
		}
		//
		if r-l+1>maxLength{
			maxLength=r-l+1
		}
		//
		r++
	}
	return maxLength
}
//暴力 超时
func longestSubarray(nums []int,limit int )int{
	l,r:=0,0
	maxLength:=0
	maxValue:=0
	minValue:=math.MaxInt32
	for r<len(nums){
		if nums[r]>maxValue{
			maxValue=nums[r]
		}
		if nums[r]<minValue{
			minValue=nums[r]
		}
		for maxValue-minValue>limit{
			if maxLength<r-l{
				maxLength=r-l
			}
			if nums[l]==minValue||nums[l]==maxValue{
				maxValue,minValue=getmaxmin(nums,l+1,r)
			}
			l++
		}
		r++
	}
	if r-l>maxLength{
		return r-l
	}
	return maxLength
}
func getmaxmin(nums []int,l,r int )(int,int){
	maxValue:=0
	minValue:=math.MaxInt32
	for i:=l;i<=r;i++{
		if nums[i]>maxValue{
			maxValue=nums[i]
		}
		if nums[i]<minValue{
			minValue=nums[i]
		}
	}
	return maxValue,minValue
}