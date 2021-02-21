package Repeat_01

import (
	"sort"
)

// 354
// important is sort
func maxEnvelopes(envelopes [][]int) int {
	if v:=len(envelopes);v<=1{
		return v
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0]<envelopes[j][0]{
			return true
		}else if envelopes[i][0]==envelopes[j][0]{
			return envelopes[i][1]>=envelopes[j][1]
		}else if envelopes[i][0]>envelopes[j][0]{
			return false
		}
		return false
	})
	//fmt.Println(envelopes)
	dp:=make([]int,len(envelopes))
	for i:=range dp{
		dp[i]=1
	}
	for i:=0;i<len(envelopes);i++{
		for j:=0;j<i;j++{
			if envelopes[i][0]>envelopes[j][0]&&envelopes[i][1]>envelopes[j][1]{
				dp[i]=maxV(dp[i],dp[j]+1)
			}
		}
	}
	max:=0
	for i:=range dp{
		if dp[i]>max{
			max=dp[i]
		}
	}
	return max
}

func maxV(a, b int )int{
	if a>b{
		return a
	}
	return b
}

// binary search
func maxEnvelopesBinarySearch(envelopes [][]int) int {
	if len(envelopes) <= 1 {
		return len(envelopes)
	}
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	tail := make([]int, 0)
	for i := 0; i < len(envelopes); i++ {
		lo, li := 0, len(tail) - 1
		for lo <= li {
			mid := lo + (li - lo)/2
			if envelopes[i][1] > tail[mid] {
				lo = mid + 1
			} else {
				li = mid - 1
			}
		}
		if lo >= len(tail) {
			tail = append(tail, envelopes[i][1])
		} else {
			tail[lo] = envelopes[i][1]
		}
	}
	return len(tail)
}
