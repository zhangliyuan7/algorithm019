package Repeat_01

import "sort"

//90
func subsetsWithDup(nums []int) (ans [][]int ){
	//不排序过不了测试用例（无序重复数字）
	sort.Ints(nums)
	var dfs func(bool,[]int,int)
	dfs = func(choosePre bool,temp []int,index int){
		//全部选择过一遍之后
		if index==len(nums){
			ans=append(ans,append([]int{},temp...))
			return
		}
		//不选当前值，此项必须放在下面判断之前，因为有可能两个连续项都不选，如果直接退出会减少可能
		dfs(false,temp,index+1)
		//前一项不选+选择当前项，会重复
		if !choosePre&&index>0&&nums[index-1]==nums[index]{
			return
		}
		//选择当前值
		temp=append(temp,nums[index])
		//选择后续，下个索引位置
		dfs(true,temp,index+1)

	}
	dfs(false,[]int{},0)
	return
}