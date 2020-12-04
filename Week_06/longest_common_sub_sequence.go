package Week_06

// 1143  二维矩阵，循环两个字符串，当相同时，等于上个i，j时刻的最长子串长度加一
// 当不同时，则取最大的i-1,j串长度或者j-1，i串长度
//if i==j {dp[i][j]=dp[i-1][j-1]+1}
// else dp[i][j]=max(dp[i-1][j],dp[i][j-1])
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	memo := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		memo[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				memo[i][j] = 1 + memo[i-1][j-1]
			} else {
				memo[i][j] = max(memo[i][j-1], memo[i-1][j])
			}
		}
	}
	return memo[m][n]
}
