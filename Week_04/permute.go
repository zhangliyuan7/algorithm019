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
func permute(nums []int)[][]int{
	var r [][]int
	for i:=0;i<len(nums);i++{
		nextr:=[]int{}
		nextr=append(nextr,nums[:i]...)
		nextr=append(nextr,nums[i+1:]...)
		tmpr:=permute(nextr)
		for i:=range tmpr{
			tmpr[i]=append(tmpr[i],nums[i])
		}
		r=append(r,tmpr...)
	}
	return r

}