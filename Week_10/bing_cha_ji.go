package Week_10


// template
//常用于找 相关性集合和圈子类问题
type BCJ struct{
	parents []int
}

func NewBCJ(n int)*BCJ{
	parents:=make([]int,n)
	//并查集常规初始化
	for i:=0;i<n;i++{
		parents[i]=i
	}
	return &BCJ{
		parents:parents,
	}
}
//search
func (bcj *BCJ)searchParents(a int )int {
	r:=a
	for bcj.parents[r]!=r{
		r=bcj.parents[r]
	}
	bcj.parents[a]=r  //剪枝 扁平化
	return r
}
//合并，将父节点统一
func  (bcj *BCJ)union(a,b int ){
	pa:=bcj.searchParents(a)
	pb:=bcj.searchParents(b)
	bcj.parents[pa]=pb
}

