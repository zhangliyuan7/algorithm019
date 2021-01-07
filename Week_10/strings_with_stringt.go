package Week_10
//115 bad solution
func Interview( s string ,t string )int{
	m:=len(s)
	n:=len(t)
	dp:=make([][]int,m+1)
	for i:=0;i<m+1;i++{
		dp[i]=make([]int,n+1)
	}
	count :=0
	for i:=1;i<m;i++{
		if lcsfunc(s[0:i],t)==n{
			count+=1
		}
	}
	return count
}

func lcsfunc(s string ,t string )int {
	m:=len(s)
	n:=len(t)
	dp:=make([][]int,m+1)
	for i:=0;i<m+1;i++{
		dp[i]=make([]int,n+1)
	}
	for i:=1;i<m+1;i++{
		for j:=1;j<n+1;j++{
			if s[i-1]==t[j-1]{
				dp[i][j]=dp[i-1][j-1]+1
			}else{
				dp[i][j]=maxlcs(dp[i][j-1],dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}
func maxlcs(a,b int )int {
	if a>b{
		return a
	}
	return   b
}