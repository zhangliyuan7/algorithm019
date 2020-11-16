package w2


//922
// 只考虑奇数位置或者偶数位置即可
func sortArrayByParityII(A []int) []int {
	var i = 0
	var j = 1
	for i<len(A)||j<len(A){
		if A[i]%2==0{
			i+=2
		}else{
			A[j],A[i]=A[i],A[j]
			j+=2
		}
	}
	return A
}

//位运算，提速
func sortArrayByParityII1(A []int) []int {
	var i = 0
	var j = 1
	for i<len(A)||j<len(A){
		if A[i]&1==0{
			i+=2
		}else{
			A[j],A[i]=A[i],A[j]
			j+=2
		}
	}
	return A
}
// loop
func sortArrayByParityII2(A []int) []int {
	i,j :=0,1
	for i<len(A){
		if A[i]&1==1{
			A[j],A[i]=A[i],A[j]
			j+=2
		}else{
			i+=2
		}
	}
	return A
}