package Repeat_01

//832
func flipAndInvertImage(A [][]int) [][]int {
	l,r:=0,len(A[0])-1
	for l<r{
		for i:=range A{
			A[i][l],A[i][r]=A[i][r]^1,A[i][l]^1
		}
		l++
		r--
	}
	// 奇数长度
	if len(A[0])&1==1{
		for i:=range A{
			A[i][l]^=1
		}
	}
	return A
}
