package Repeat_01

// 1579
func MaxNumEdgesToRemove(n int, edges [][]int) int {
	ufA := newUnionFindy(n)
	ufB := newUnionFindy(n)
	ans := len(edges)
	for i := 0; i < len(edges); i++ {
		x, y := edges[i][1]-1, edges[i][2]-1
		if edges[i][0] == 3 && (!ufA.inSameSet(x, y) || !ufB.inSameSet(x, y)) {
			ufA.union(x, y)
			ufB.union(x, y)
			ans--
		}
	}
	for i := 0; i < len(edges); i++ {
		x, y := edges[i][1]-1, edges[i][2]-1
		if edges[i][0] == 2 && ufB.union(x, y) {
			ans--
		}
		if edges[i][0] == 1 && ufA.union(x, y) {
			ans--
		}
	}
	if ufA.count > 1 || ufB.count > 1 {
		return -1
	}
	return ans
}

type unionFindy struct {
	parents []int
	size    []int
	count   int
}

func (uf *unionFindy) inSameSet(x, y int) bool {
	return uf.find(x) == uf.find(y)
}

func newUnionFindy(n int) unionFindy {
	uf := unionFindy{parents: make([]int, n), size: make([]int, n), count: n}
	for i := 0; i < n; i++ {
		uf.parents[i] = i
		uf.size[i] = 1
	}
	return uf
}
func (uf *unionFindy) find(x int) int {
	fx := uf.parents[x]
	if fx != x {
		uf.parents[x] = uf.find(fx)
	}
	return uf.parents[x]
}

func (uf *unionFindy) union(x, y int) bool {
	fx := uf.find(x)
	fy := uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.parents[fy] = fx
	uf.parents[y] = fx
	uf.size[fx] += uf.size[fy]
	uf.count--
	return true
}
