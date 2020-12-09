package Week_06

//bottom - up
func minimumTotal(triangle [][]int) int {
	dp:=triangle
	for i:=len(triangle)-2;i>=0;i--{
		for j:=0;j<len(triangle[i]);j++{
			dp[i][j]=min(dp[i+1][j],dp[i+1][j+1])+dp[i][j]
		}
	}
	return dp[0][0]
}

// top - down
func minimumTotal2(triangle [][]int )int {
	return recursion_minimum_total(triangle,0,0)
}

func recursion_minimum_total(triangle [][]int ,level int ,index int )int {
	if level== len(triangle){
		return triangle[level][index]
	}
	l:=recursion_minimum_total(triangle,level+1,index)
	r:=recursion_minimum_total(triangle,level+1,index+1)
	return min(l,r)+triangle[level][index]
}