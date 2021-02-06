package Repeat_01

import (
	"math"
)

// 1423 sliding window
func maxScore(cardPoints []int, k int) int {
	windownSize:=len(cardPoints)-k
	minPoints:=math.MaxInt32
	tmpPoints:=0
	l:=0
	sum:=0
	for r:=0;r<len(cardPoints);r++{
		sum+=cardPoints[r]
		if r-l+1<=windownSize{
			tmpPoints+=cardPoints[r]
		}else{
			minPoints=min(minPoints,tmpPoints)
			tmpPoints=tmpPoints-cardPoints[l]+cardPoints[r]
			l++
		}
	}
	//fmt.Println(sum)
	minPoints=min(minPoints,tmpPoints)
	return sum-minPoints
}

func min(a,b int )int{
	if a<b{
		return a
	}
	return b
}

//dp
func maxScoreII(cardPoints []int,k int )int{
	left:= make([]int,k+1)
	right:= make([]int,k+1)
	n:=len(cardPoints)
	for i:=1;i<k+1;i++ {
		left[i]=left[i-1]+cardPoints[i-1]
	}
	maxPoints:=left[k]
	for i:=1;i<k+1;i++{
		right[i]=right[i-1]+cardPoints[n-i]
		maxPoints=max(maxPoints,right[i]+left[k-i])
	}
	return maxPoints
}