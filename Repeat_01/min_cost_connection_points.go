package Repeat_01

import (
	"math"
	"sort"
)

// bingchaji 1584
type UnionFind struct{
	parents []int
	rank []int
}

func newUnionFind(n int )*UnionFind{
	uf:=UnionFind{}
	uf.parents=make([]int,n)
	uf.rank=make([]int,n)
	for i:=0;i<n;i++{
		uf.parents[i]=i
		uf.rank[i]=1
	}
	return &uf
}

func (uf *UnionFind)Find(x int )int{
	if  uf.parents[x]!=x{
		uf.parents[x] =uf.Find(uf.parents[x])
	}
	return uf.parents[x]
}
func (uf *UnionFind)Union(x,y int )bool{
	fx:=uf.Find(x)
	fy:=uf.Find(y)
	if fx==fy {return false }
	if uf.rank[fx]<uf.rank[fy]{
		fx,fy=fy,fx
	}
	uf.rank[fx]+=uf.rank[fy]
	uf.parents[fy]=fx
	return true
}

func distance(p,q []int )int{
	return int(math.Abs(float64(p[0]-q[0]))+math.Abs(float64(p[1]-q[1])))
}

type edge struct{
	start int
	end int
	dis int
}

func minCostConnectPoints(points [][]int)(ans int) {
	n:=len(points)
	uf:=newUnionFind(n)
	edges:=[]edge{}
	for i:=0;i<len(points);i++{
		for j:=i;j<len(points);j++{
			edges=append(edges,edge{i,j,distance(points[i],points[j])})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].dis<edges[j].dis
	})
	left:=n-1
	for i:=0;i<len(edges);i++{
		if uf.Union(edges[i].start,edges[i].end){
			ans+=edges[i].dis
			left--
			if left==0{
				break
			}
		}
	}
	return ans
}