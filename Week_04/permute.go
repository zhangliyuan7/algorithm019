package Week_04

func Permute(nums []int) [][]int {
	var r [][]int
	if len(nums)==1{
		return [][]int{nums}
	}
	for i:=0;i<len(nums);i++{
		var others []int
		others=append(others,nums[:i]...)
		others=append(others,nums[i+1:]...)
		tmpR:=Permute(others)
		for x:=0;x<len(tmpR);x++{
			tmpR[x]=append(tmpR[x],nums[i])
		}
		r=append(r,tmpR...)
	}
	return r
}