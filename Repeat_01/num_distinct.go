package Repeat_01

// 115
func numDistinct(s string, t string) int {
	ls:=len(s)
	lt:=len(t)
	dp:=make([][]int,ls+1)
	for i:=range dp{
		dp[i]=make([]int,lt+1)
		dp[i][0]=1
	}
	for i:=1;i<ls+1;i++{
		for j:=1;j<lt+1;j++{
			if s[i-1]==t[j-1]{
				dp[i][j]=dp[i-1][j-1]+dp[i-1][j]
			}else{
				dp[i][j]=dp[i-1][j]
			}
		}
	}
	return dp[ls][lt]
}
