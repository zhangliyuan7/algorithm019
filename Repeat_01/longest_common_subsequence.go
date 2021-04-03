package Repeat_01


//1143
func longestCommonSubsequence(text1 string, text2 string) int {
	ln1:=len(text1)
	ln2:=len(text2)
	dp:=make([][]int,ln1)
	for i:=range dp{
		dp[i]=make([]int,ln2)
	}
	for i:=1;i<ln1+1;i++{
		for j:=1;j<ln2+1;j++{
			if text1[i-1]==text2[j-1]{
				dp[i][j]=dp[i-1][j-1]+1
			}else{
				dp[i][j]=max(dp[i-1][j],dp[i][j-1])
			}
		}
	}
	return dp[ln1][ln2]
}
