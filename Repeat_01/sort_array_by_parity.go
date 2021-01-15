package Repeat_01

// 905
func sortArrayByParity(A []int) []int {
	right:=len(A)-1
	for i:=0;i<right;i++{
		if A[i]&1==1{
			A[i],A[right]=A[right],A[i]
			right--
			i--
		}
	}
	return A
}