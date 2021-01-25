package Repeat_01

// 959
// 生成矩阵思考较为困难，其余dfs简单粗暴的修改
func regionsBySlashes(grid []string) int {
	matrix:=make([][]int,3*len(grid),3*len(grid[0]))

	for i:=range matrix{
		matrix[i]=make([]int,3*len(grid[0]))
	}
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j]==' '{
				continue
			}
			if grid[i][j]=='/'{
				matrix[i*3+2][j*3]=1
				matrix[i*3+1][j*3+1]=1
				matrix[i*3][j*3+2]=1
			}
			if grid[i][j]=='\\'{
				matrix[i*3][j*3]=1
				matrix[i*3+1][j*3+1]=1
				matrix[i*3+2][j*3+2]=1
			}
		}
	}
	//fmt.Println(matrix)
	island:=0
	for i:=0;i<len(matrix);i++{
		for j:=0;j<len(matrix[0]);j++{
			if matrix[i][j]==0{
				island+=1
				dfsSetOne(i,j,matrix)
			}
		}
	}
	return island
}

func dfsSetOne(i,j int,matrix [][]int) {
	if i<0||j<0||i>=len(matrix)||j>=len(matrix[0])||matrix[i][j]==1{
		return
	}
	matrix[i][j]=1
	dfsSetOne(i+1,j,matrix)
	dfsSetOne(i-1,j,matrix)
	dfsSetOne(i,j+1,matrix)
	dfsSetOne(i,j-1,matrix)
}


