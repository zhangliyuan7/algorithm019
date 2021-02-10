package Repeat_01

// 992
// timeout , O(n*n)
func subarraysWithKDistinct(A []int, K int) int {
	la:=len(A)

	if la < K {
		return 0
	}
	ans := 0
	for l:=0;l<la;l++ {
		m:=make(map[int]int)
		for r:=l;r<la;r++{
			m[A[r]]+=1
			if len(m)==K{
				ans++
			}
			if len(m)>K{
				break
			}
		}
	}
	return ans
}

func mostKSubArraysWithDistinctNumber(A []int,K int )int{
	ret:=0
	l,r:=0,0
	la:=len(A)
	distinct:=0
	m:=make(map[int]int)
	for ;r<la;r++{
		if m[A[r]]==0{
			distinct+=1
		}
		m[A[r]]+=1
		for distinct>K{
			m[A[l]]-=1
			if m[A[l]]==0{
				distinct-=1
			}
			l++
		}
		ret+=r-l+1
	}
	return ret
}

//恰好k个不同元素数组个数 = 最多k个不同元素的子数组个数  - 最多k-1个不同元素的子数组个数
func subarraysWithKDistinctSlidingWindow(A []int, K int) int {
	if len(A)<K{
		return 0
	}
	return mostKSubArraysWithDistinctNumber(A,K)-mostKSubArraysWithDistinctNumber(A,K-1)
}
