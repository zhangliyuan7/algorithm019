package w3

import "math"

// ez - 1030
//暴力解决，list索引存储距离，保证顺序
//比bfs(44ms)快多了。。。20ms
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	var disList = [200][][]int{}
	for c:=0;c<C;c++{
		for r:=0;r<R;r++{
			dis:=int(math.Abs(float64(r0-r))+math.Abs(float64(c0-c)))
			disList[dis]=append(disList[dis],[]int{r,c})
		}
	}
	var r [][]int
	for _,v:=range disList{
		if len(v)==0{
			continue
		}
		r=append(r,v...)
	}
	return r
}

//bfs
//慢的一笔 ，44ms
func AllCellsDistOrder1(R int, C int, r0 int, c0 int) [][]int {
	var res [][]int
	res=append(res,[]int{r0,c0})
	var visited =make(map[[2]int]bool)
	//visit top,bottom,left,right
	visited[[2]int{r0,c0}]=true
 	next:=[][2]int{{r0,c0}}
	for len(next)!=0 {
		r,c:=next[0][0],next[0][1]
		//up
		if c+1<C&&visited[[2]int{r,c+1}]!=true{
			visited[[2]int{r,c+1}]=true
			res=append(res,[]int{r,c+1})
			next=append(next,[2]int{r,c+1})
		}
		//down
		if c-1>=0&&visited[[2]int{r,c-1}]!=true{
			visited[[2]int{r,c-1}]=true
			res=append(res,[]int{r,c-1})
			next=append(next,[2]int{r,c-1})
		}
		//right
		if r+1<R&&visited[[2]int{r+1,c}]!=true{
			visited[[2]int{r+1,c}]=true
			res=append(res,[]int{r+1,c})
			next=append(next,[2]int{r+1,c})
		}
		//left
		if r-1>=0&&visited[[2]int{r-1,c}]!=true{
			visited[[2]int{r-1,c}]=true
			res=append(res,[]int{r-1,c})
			next=append(next,[2]int{r-1,c})
		}
		next=next[1:]
	}
	return res
}

func allCellsDistOrderBucketSort(R int, C int, r0 int, c0 int) [][]int {
	var bucket = [200][][]int{}
	var res =[][]int{}
	for r:=0;r<R;r++{
		for c:=0;c<C;c++{
			dis:=distance(r,c,r0,c0)
			bucket[dis]=append(bucket[dis],[]int{r,c})
		}
	}
	for  i:=0;i<len(bucket);i++{
		if len(bucket[i])==0{
			continue
		}
		res=append(bucket[i])
	}
	return res
}

func distance(r,c,r0,c0 int )int{
	return int(math.Abs(float64(r-r0))+math.Abs(float64(c-c0)))
}

func  allCellsDistOrderBfs(R int, C int, r0 int, c0 int) [][]int {
	var res [][]int
	visited:=make(map[[2]int]bool)
	queue:=[][2]int{{r0,c0}}
	for len(queue)!=0{
		r,c:=queue[0][0],queue[0][1]
		queue=queue[1:]
		if r+1<R&&!visited[[2]int{r+1,c}]{
			res=append(res,[]int{r+1,c})
			queue=append(queue,[2]int{r+1,c})
			visited[[2]int{r+1,c}]=true
		}
		if c+1<C&&!visited[[2]int{r,c+1}]{
			res=append(res,[]int{r,c+1})
			queue=append(queue,[2]int{r,c+1})
			visited[[2]int{r,c+1}]=true
		}
		if r-1>=0&&!visited[[2]int{r-1,c}]{
			res=append(res,[]int{r-1,c})
			queue=append(queue,[2]int{r-1,c})
			visited[[2]int{r-1,c}]=true
		}
		if c-1 >=0&&!visited[[2]int{r,c-1}]{
			res=append(res,[]int{r,c-1})
			queue=append(queue,[2]int{r,c-1})
			visited[[2]int{r,c-1}]=true
		}
	}
	return res
}