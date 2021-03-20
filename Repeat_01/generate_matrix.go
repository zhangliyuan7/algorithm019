package Repeat_01


//59 模拟法
func generateMatrix(n int) [][]int {
	directx:=[]int{0,1,0,-1}
	directy:=[]int{1,0,-1,0}
	matrix:=make([][]int,n)
	for i:=range matrix{
		// 因为是 [1,n*n] 没有0值，所以不用赋初值 0即可
		matrix[i]=make([]int,n)
	}
	row,col:=0,0
	direct:=0
	for i:=1;i<=n*n;i++{
		matrix[row][col]=i
		r:=row+directx[direct]
		c:=col+directy[direct]
		// 外圈只需要处理三个角，内圈处理判断已经写入
		if c==-1||c==n||r==n||(r<n && c<n && matrix[r][c]!=0) {
			direct+=1
			if direct==4{
				direct = 0
			}
		}
		row+=directx[direct]
		col+=directy[direct]
	}
	return matrix
}