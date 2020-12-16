package Week_07

//547 并查集
func FindCircleNum(M [][]int) int {
	var parents =make([]int,len(M))
	for i:=range M{
		parents[i] = i
	}
	count:=0
	for i:=0;i<len(M);i++{
		for j:=0;j<len(M);j++{
			if M[i][j]==1&&i!=j{
				union(&parents,i,j )
			}
		}
	}
	for i:=range parents{
		if parents[i] ==i {
			count+=1
		}
	}
	return count
}

func union(p *[]int ,i , j int)  {
	p1:=searchParents(p,i)
	p2:=searchParents(p,j)
	(*p)[p1]=p2
}

func searchParents(p *[]int,i int )int {
	r:=i
	for (*p)[r]!=r{
		r=(*p)[r]
	}
	(*p)[i]=r
	return r
}