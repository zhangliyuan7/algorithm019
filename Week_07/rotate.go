package Week_07


// 48
// i,j 经过右旋转一次，坐标变为 j,n-i-1
// 赋值的时候 ，按照旋转方向赋值， 比如 0，0的位置 由 n，0 的位置取值赋给他
func rotate(matrix [][]int )[][]int{
	n:=len(matrix)
	for i:=0;i<n/2;i++{
		for j:=0;j<n/2;j++{
			matrix[i][j],matrix[n-j-1][i],matrix[n-i-1][n-j-1],matrix[j][n-i-1]=
				matrix[n-j-1][i],matrix[n-i-1][n-j-1],matrix[j][n-i-1],matrix[i][j]
		}
	}
	return matrix
}