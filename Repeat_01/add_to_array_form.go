package Repeat_01

// 989
func AddToArrayForm(A []int, K int) []int {
	l:=len(A)
	front:=0
	for i:=l-1;i>=0;i--{
		k:=K%10
		K=K/10
		s:=A[i]+k+front
		front=s/10
		A[i]=s%10
	}
	leftK:=[]int{}
	for K!=0||front!=0{
		s:=K%10+front
		K=K/10
		front=s/10
		leftK=append(leftK,s%10)
	}
	r:=[]int{}
	for i:=len(leftK)-1;i>=0;i--{
		r=append(r,leftK[i])
	}
	r=append(r,A...)
	return r
}
