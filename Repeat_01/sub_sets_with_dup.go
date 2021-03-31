package Repeat_01

import "sort"

//90
func subsetsWithDup(nums []int) (ans [][]int ){
	sort.Ints(nums)
	var dfs func(bool,[]int,int)
	dfs = func(choosePre bool,temp []int,index int){
		if index==len(nums){
			ans=append(ans,append([]int{},temp...))
			return
		}
		dfs(false,temp,index+1)
		if !choosePre&&index>0&&nums[index-1]==nums[index]{
			return
		}
		temp=append(temp,nums[index])
		dfs(true,temp,index+1)
	}
	dfs(false,[]int{},0)
	return
}