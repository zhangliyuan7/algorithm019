package Repeat_01

//78
// every step choose or not choose
//
func subsets(nums []int) (ans [][]int) {
	var dfs func([]int,int)
	dfs=func(temp []int,index int){
		//如果整个nums都经历了一次选择，存储结果
		if index==len(nums){
			ans=append(ans,append([]int{},temp...))
			return
		}
		//是否选择当前索引数字 选择
		temp=append(temp,nums[index])
		//到下个索引位置
		dfs(temp,index+1)
		//不选择
		temp=temp[:len(temp)-1]
		//到下个索引位置
		dfs(temp,index+1)
	}
	dfs([]int{},0)
	return
}

