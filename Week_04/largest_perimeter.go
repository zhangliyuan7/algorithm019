package Week_04

import (
	"sort"
)
func largestPerimeter(A []int) int {
	if len(A)<3{
		return 0
	}
	sort.Ints(A)
	for i:=len(A)-1;i>=2;i--{
		if A[i-2]+A[i-1]-A[i]>0{
			return  A[i-2]+A[i-1]+A[i]
		}
	}
	return 0
}

func LargestPerimeter(A []int) int {
	 if len(A)<3{
	 	return 0
	 }
	sort.Ints(A)
	 for i:=len(A)-1;i>=2;i--{
	 	if valid(A[i],A[i-1],A[i-2]){
	 		return A[i]+A[i-1]+A[i-2]
		}
	 }
	 return 0
}

func valid(a,b,c int )bool{
	if a+b<=c{
		return false
	}
	if a+c<=b{
		return false
	}
	if c+b<=a{
		return false
	}
	return true
}