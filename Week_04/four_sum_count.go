package Week_04

func FourSumCount(A []int, B []int, C []int, D []int) int {
	ab := make(map[int]int)
	var r int
	length:=len(A)
	for i:=0;i<length;i++{
		for j:=0;j<length;j++{
			ab[A[i]+B[j]]++
		}
	}
	for i:=0;i<length;i++{
		for j:=0;j<length;j++{
				r+=ab[-1*(C[i]+D[j])]
		}
	}
	return r
}
