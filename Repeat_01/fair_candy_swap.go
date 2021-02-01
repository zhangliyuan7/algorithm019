package Repeat_01

// 888
func fairCandySwap(A []int, B []int) []int {
	sumA:=0
	sumB:=0
	setA:=make(map[int]int)
	for i:=range A{
		sumA+=A[i]
		setA[A[i]]=0
	}
	for j:=range B{
		sumB+=B[j]
	}
	dis:=(sumB-sumA)/2
	for i:=range B{
		if _,ok:=setA[B[i]-dis];ok{
			return []int{B[i]-dis,B[i]}
		}
	}
	return nil
}
