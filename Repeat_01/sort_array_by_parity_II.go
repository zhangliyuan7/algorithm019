package Repeat_01

//922
func SortArrayByParityII(A []int) []int {
	i:=0
	j:=1
	for i<len(A)&&j<len(A){
		if A[i]&1==0{
			i+=2
		}else{
			A[j],A[i]=A[i],A[j]
			j+=2
		}
	}
	return A
}