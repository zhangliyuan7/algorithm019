package Week_09

import "math"

// 72
func minDistance(word1 string, word2 string) int {
	m:=len(word1)+1
	n:=len(word2)+1
	dp:=make([][]int,m)
	for i:=range dp{
		dp[i]=make([]int,n)
	}
	for i:=0;i<m;i++{
		dp[i][0]=i
	}
	for i:=0;i<n;i++{
		dp[0][i]=i
	}
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			// 否则少一次word的子母判断
			if word1[i-1]==word2[j-1]{
				dp[i][j]=dp[i-1][j-1]
			}else{
				dp[i][j]=minthree(dp[i-1][j-1]+1,dp[i-1][j]+1,dp[i][j-1]+1)
			}
		}
	}
	return dp[m-1][n-1]
}

func minthree(a,b,c int )int{
	candidates:=[]int{a,b,c}
	min:=math.MaxInt32
	for i:=0;i<3;i++{
		if  candidates[i]< min{
			min=candidates[i]
		}
	}
	return min
}

func editDistance(s1 ,s2 string )int {
	m:=len(s1)
	n:=len(s2)
	dp:=make([][]int,m+1)
	for i:=0;i<m+1;i++{
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
			if s1[i-1]==s2[j-1]{
				dp[i][j]=dp[i-1][j-1]
			}else{
				dp[i][j]=minThree(dp[i-1][j-1],dp[i-1][j],dp[i][j-1])+1
			}
		}
	}
	return dp[m][n]
}
func minThree(a,b,c int )int {
	tmp:=[]int{a,b,c}
	min:=a
	for i:=0;i<len(tmp);i++{
		if min>tmp[i]{
			min=tmp[i]
		}
	}
	return min
}