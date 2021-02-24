package Repeat_01


//64
func minPathSum(grid [][]int) int {
	m,n:=len(grid),len(grid[0])
	dp:=make([][]int,m)
	for i:=0;i<m;i++{
		dp[i]=make([]int,n)
	}
	dp[0][0]=grid[0][0]
	for i:=1;i<n;i++{
		dp[0][i]=dp[0][i-1]+grid[0][i]
	}
	for i:=1;i<m;i++{
		dp[i][0]=dp[i-1][0]+grid[i][0]
	}
	//fmt.Println(dp)
	for i:=1;i<m;i++{
		for j:=1;j<n;j++{
			dp[i][j]=min(dp[i][j-1],dp[i-1][j])+grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
