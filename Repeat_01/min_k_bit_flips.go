package Repeat_01

// 995
// 暴力法超时
func minKBitFlips(A []int, K int) int {
	l:=0
	r:=0
	count:=0
	for r<len(A) {
		if r<K-1{
			r++
			continue
		}
		if A[l]==0{
			count++
			for i:=l;i<=r;i++{
				if A[i]!=1{
					A[i]=1
				}else{
					A[i]=0
				}
			}
			l++
		}
		r++
	}
	for i:=l;i<len(A);i++{
		if A[i]==0{
			return -1
		}
	}
	return count
}

//差分
func minKBitFlipsB(A []int, K int) (ans int) {
	n := len(A)
	diff := make([]int, n+1)
	revCnt := 0
	for i, v := range A {
		revCnt += diff[i]
		if (v+revCnt)%2 == 0 {
			if i+K > n {
				return -1
			}
			ans++
			revCnt++
			diff[i+K]--
		}
	}
	return
}