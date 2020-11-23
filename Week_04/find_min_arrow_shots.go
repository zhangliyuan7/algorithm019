package Week_04

import (
	"sort"
)

func FindMinArrowShots(points [][]int) int {
	if len(points)<=1{
		return len(points)
	}
	// 右边界排序
	sort.Slice(points, func(i, j int) bool {
		if (points[i][1]<points[j][1]){
			return true
		}
		return false
	})
	// fmt.Println(points)
	var count int = 1
	var pos int =points[0][1]
	for i:=1;i<len(points);i++{
		// 起始点>前面气球的结束点，说明重叠结束
		// 修改下一个射箭位置为不重叠的下个气球的右边界
		if points[i][0]>pos{
			count+=1
			pos=points[i][1]
		}
	}
	return count
}