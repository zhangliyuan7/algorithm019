package Week_10

//1143
func lcs(w1,w2 string)int {
	m:=len(w1)
	n:=len(w2)
	dp:=make([][]int,m+1)
	for i:=range dp{
		dp[i]=make([]int,n+1)
	}
	for i:=1;i<m+1;i++{
		for j:=1;j<n+1;j++{
			if w1[i-1]==w2[j-1]{
				dp[i][j]=dp[i-1][j-1]+1
			}else{
				dp[i][j]=maxValue(dp[i-1][j],dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
func maxValue(a,b int )int {
	if a>b{
		return a
	}
	return b
}