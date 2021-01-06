package Week_10

// 399
func CalcEquation(equations [][]string ,values []float64,queries [][]string)[]float64{
	id :=make(map[string]int)
	for _,eq:=range equations{
		fir,sec:=eq[0],eq[1]
		if _,has:=id[fir];!has{
			id[fir]=len(id)
		}
		if _,has:=id[sec];!has{
			id[sec]=len(id)
		}
	}
	bcj:=newbcj(len(id))
	for i,eq:=range equations{
		bcj.union(id[eq[0]],id[eq[1]],values[i])
	}
	ans:=make([]float64,len(queries))
	for i,q:=range queries{
		start,hasX:=id[q[0]]
		end,hasY:=id[q[1]]
		if hasX && hasY&& bcj.findParents(start) == bcj.findParents(end){
			ans[i]=bcj.w[start]/bcj.w[end]
		}else{
			ans[i]=-1
		}
	}
	return ans
}

type bcj struct{
	fa []int
	w []float64
}
func newbcj(l int )*bcj{
	bcj:=&bcj{
		make([]int,l),
		make([]float64,l),
	}
	for i:=range bcj.fa{
		bcj.fa[i]=i
		bcj.w[i]=1
	}
	return bcj
}
func (b *bcj)findParents(x int )int {
	if b.fa[x]!=x{
		ff:=b.findParents(b.fa[x])
		b.w[x]*=b.w[b.fa[x]]
		b.fa[x]=ff
	}
	return b.fa[x]
}

func (b *bcj)union(from,to int,value float64 ){
	ff:=b.findParents(from)
	ft:=b.findParents(to)
	b.w[ff]=value*b.w[to]/b.w[from]
	b.fa[ff]=ft
}

// official solution
func calcEquationX(equations [][]string, values []float64, queries [][]string) []float64 {
	// 给方程组中的每个变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	fa := make([]int, len(id))
	w := make([]float64, len(id))
	for i := range fa {
		fa[i] = i
		w[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			f := find(fa[x])
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}
	merge := func(from, to int, val float64) {
		fFrom, fTo := find(from), find(to)
		w[fFrom] = val * w[to] / w[from]
		fa[fFrom] = fTo
	}

	for i, eq := range equations {
		merge(id[eq[0]], id[eq[1]], values[i])
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if hasS && hasE && find(start) == find(end) {
			ans[i] = w[start] / w[end]
		} else {
			ans[i] = -1
		}
	}
	return ans
}
