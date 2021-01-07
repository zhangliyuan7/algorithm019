package Week_10


//115
func numDistinct(s string, t string) int {
	m:=len(s)
	n:=len(t)
	dp:=make([][]int,n+1)
	for i:=range dp{
		dp[i]=make([]int,m+1)
	}
	for i:=range dp[0]{
		//空字符串 是所有长度的s的子序列
		dp[0][i]=1
	}
	for i:=0;i<n+1;i++{
		for j:=1;j<m+1;j++{
			if s[j-1]==t[i-1]{
				// 是否选择s对应j位置的字母，选择和不选择的结果要相加，因为是不同子序列
				dp[i][j]=dp[i][j-1]+dp[i-1][j-1]
			}else{
				// 因为不同，所以就是唯一子序列，不会出现两个相同重复子序列，所以只取其中一个值
				dp[i][j]=dp[i][j-1]
			}
		}
	}
	return dp[n][m]
}