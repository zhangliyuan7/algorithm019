package Week_02

import (
	"sort"
)

// 利用计数 有序的计数 做写入控制

func relativeSortArray1(arr1 []int, arr2 []int) []int {
	var r []int
	var upper int
	for i:=range arr1{
		if arr1[i]>upper{
			upper = arr1[i]
		}
	}
	var fre =make([]int,upper+1)
	for i:=0;i<len(arr1);i++{
		fre[arr1[i]]++
	}
	for i:=0;i<len(arr2);i++ {
		for ; fre[arr2[i]] > 0; fre[arr2[i]]-- {
			r = append(r, arr2[i])
		}
	}
	//fre列表元素是次数，元素索引是值，索引顺序为元素大小顺序
	for i:=0;i<len(fre);i++{
		for ; fre[i]>0;fre[i]--{
			r=append(r,i)
		}
	}
	return r
}

// yong 自定义sort解决，注意理解题目，是相对顺序不变，并不是把所有的arr2的元素全放在前面
func relativeSortArray2(arr1 []int, arr2 []int) []int {
	var m =make(map[int]int)
	for i,v:=range arr2{
		m[v]=i
	}
	sort.Slice(arr1,func(i int ,j int )bool{
		x,y:=arr1[i],arr1[j]
		idA,hasA:= m[x]
		idB,hasB:=m[y]
		if hasA&&hasB{
			return idA<idB
		}
		// A存在则当作较小的放在前面，因为存在是true 不存在是false return hasA即可
		if hasA||hasB{
			return hasA
		}
		return x<y
	})
	return arr1
}
