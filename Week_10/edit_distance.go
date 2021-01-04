package Week_10

import "math"

// 72
func editDistance(w1,w2 string )int {
	m:=len(w1)
	n:=len(w2)
	dp:=make([][]int,m+1)
	for i:=range dp{
		dp[i]=make([]int,n+1)
	}
	for i:=0;i<m+1;i++{
		dp[i][0]=i
	}
	for i:=0;i<n+1;i++{
		dp[0][i]=i
	}
	for i:=1;i<m+1;i++{
		for j:=1;j<n+1;j++{
			if w1[i-1]==w2[j-1]{
				dp[i][j]=dp[i-1][j-1]
			}else{
				dp[i][j]=minValue(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
			}
		}
	}
	return dp[m][n]
}
func minValue(a,b,c int )int {
	cand:=[]int{a,b,c}
	min:=math.MaxInt32
	for i:=range cand{
		if cand[i]<min{
			min=cand[i]
		}
	}
	return min
}