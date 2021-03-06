package Repeat_01

// 959
//
type unionFind struct {
	parent, size []int
	setCount     int // 当前连通分量数目
}

func newUnionFindx(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent, size, n}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
	uf.setCount--
}

//func regionsBySlashesBcj(grid []string) int {
	//n := len(grid)
	//uf := newUnionFindx(4 * n * n)
	//for i := 0; i < n; i++ {
	//	for j := 0; j < n; j++ {
	//		idx := i*n + j
	//		if i < n-1 {
	//			bottom := idx + n
	//			uf.union(idx*4+2, bottom*4)
	//		}
	//		if j < n-1 {
	//			right := idx + 1
	//			uf.union(idx*4+1, right*4+3)
	//		}
	//		if grid[i][j] == '/' {
	//			uf.union(idx*4, idx*4+3)
	//			uf.union(idx*4+1, idx*4+2)
	//		} else if grid[i][j] == '\\' {
	//			uf.union(idx*4, idx*4+1)
	//			uf.union(idx*4+2, idx*4+3)
	//		} else {
	//			uf.union(idx*4, idx*4+1)
	//			uf.union(idx*4+1, idx*4+2)
	//			uf.union(idx*4+2, idx*4+3)
	//		}
	//	}
	//}
	//return uf.setCount
//}

//    0
//3       1  => / {0,3}{1,2}  \\ {2,3}{0,1}  ' ' {0,1,2,3}
//    2
func regionsBySlashesBcj(grid []string) int {
	l:=len(grid)
	uf:=newUnionFindx(4*l*l)
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			idx := i*l + j
			if i<l-1{
				btm:=idx+l
				// btm = (i+1)*l +j
				uf.union(idx*4+2,btm*4)
				//uf.union(idx*4+2,btm*4+0)
			}
			if j<l-1{
				right:=idx+1
				uf.union(idx*4+1,right*4+3)
			}
			if grid[i][j]==' '{
				uf.union(idx*4,idx*4+1)
				uf.union(idx*4+1,idx*4+2)
				uf.union(idx*4+3,idx*4+2)
			}
			if grid[i][j]=='\\'{
				uf.union(idx*4,idx*4+1)
				uf.union(idx*4+2,idx*4+3)
			}
			if grid[i][j]=='/'{
				uf.union(idx*4,idx*4+3)
				uf.union(idx*4+1,idx*4+2)

			}
		}
	}
	return uf.setCount
}
