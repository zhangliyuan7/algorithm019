package Repeat_01

import "sort"

//88
func mergeI(nums1 []int, m int, nums2 []int, n int)  {
	if n==0{
		return
	}
	a:=[]int{}
	index1:=0
	index2:=0
	for index1<m&&index2<n{
		if nums1[index1]<nums2[index2]{
			a=append(a,nums1[index1])
			index1++
		}else{
			a=append(a,nums2[index2])
			index2++
		}
	}
	if index1<m{
		a=append(a ,nums1[index1:m]...)
	}else if index2<n{
		a=append(a,nums2[index2:n]...)
	}
	copy(nums1, a )
}

func mergeII(nums1 []int, m int, nums2 []int, n int){
	copy(nums1[m:],nums2)
	sort.Ints(nums1)
}