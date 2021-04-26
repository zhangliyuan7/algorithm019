package Week_07

//200 并查集
func numIslands(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	if row == 0 || col == 0 {
		return 0
	}
	var count int
	var p = make([]int, row*col)
	for i := range grid {
		for j := range grid[0] {
			p[i*col+j] = i*col + j
			if grid[i][j]=='1'{
				count++
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				if j-1>=0&&grid[i][j-1] == '1' {
					count=unionisland(&p, i*col+(j-1), i*col+j,count)
				}
				if j+1<col&&grid[i][j+1] == '1' {
					count=unionisland(&p, i*col+(j+1), i*col+j,count)
				}
				if i+1<row&&grid[i+1][j] == '1' {
					count=unionisland(&p, (i+1)*col+j, i*col+j,count)
				}
				if i-1>=0&&grid[i-1][j] == '1' {
					count=unionisland(&p, (i-1)*col+j, i*col+j,count)
				}
			}
			grid[i][j]='0'
		}
	}
	return count
}

func unionisland(p *[]int, i, j int,count int ) int {
	pi := parentsofisland(p, i)
	pj := parentsofisland(p, j)
	if pi==pj{
		return count
	}
	(*p)[pi] = pj
	return count-1
}

func parentsofisland(p *[]int, i int) int {
	tmpi := i
	for (*p)[i] != i {
		i = (*p)[i]
	}
	(*p)[tmpi] = i //压缩路径
	return i
}

// dfs
func numIslandsdfs(grid [][]byte) int {
	row:=len(grid)
	col:=len(grid[0])
	var count int
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			if grid[i][j]=='1'{
				count+=1
				visit(i,j,&grid)
			}
		}
	}
	return count
}
func visit(i,j int ,g *[][]byte){
	if i<0||j<0||i>len(*g)-1||j>len((*g)[0])-1{
		return
	}
	if (*g)[i][j]=='0'{
		return
	}
	dx:=[]int{0,-1,0,1}
	dy:=[]int{1,0,-1,0}
	(*g)[i][j]='0'
	for d:=0;d<4;d++{
		visit(i+dx[d],j+dy[d],g)
	}
}