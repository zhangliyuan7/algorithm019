package Week_08

import "sort"

// 455 分发饼干
func FindContentChildren(g []int, s []int) int {
	if len(g)==0||len(s)==0{
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	i,j:=0,0
	left,count:=0,0
	for i<len(g)&&j<len(s){
		left+=s[j]
		if left>=g[i]{
			i++
			j++
			left=0
			count++
		}else{
			j++
		}
	}
	return count
}