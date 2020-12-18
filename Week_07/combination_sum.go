package Week_07

// 39
func CombinationSum(candidates []int, target int) [][]int {
	var r [][]int
	for i:=0;i<len(candidates);i++{
		if target>candidates[i]{
			tmpi:=candidates[i]
			for _,child_list:=range CombinationSum(candidates[i:],target-tmpi){
				child_list=append([]int{tmpi},child_list...)
				r=append(r,child_list)
			}
		}
		if target == candidates[i]{
			r=append(r,[]int{candidates[i]})
		}
	}
	return r
}