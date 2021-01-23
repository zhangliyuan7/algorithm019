package Repeat_01

// 1319
func makeConnected(n int, connections [][]int) int {
	// 排除线数量不够（小于n-1）的情况
	if len(connections)<n-1{
		return -1
	}
	// 到这里线的总数一定足够
	parents:=make([]int,n)
	for i:=0;i<n;i++{
		parents[i]=i
	}
	var findp func(a int )int
	findp=func(a int )int{
		if a!=parents[a]{
			parents[a]=findp(parents[a])
		}
		return parents[a]
	}
	var union func (a,b int )
	union =func (a,b int ){
		fa:=findp(a)
		fb:=findp(b)
		if fa==fb{
			return
		}
		parents[fb]=fa
		parents[b]=fa
	}
	for i:=0;i<len(connections);i++{
		union(connections[i][0], connections[i][1])
	}
	// 合并结束，parents中剩余多少没合并的局域网，最少的链接方式 就是其数量减一（点-1=线）
	disconnectionCount:=0
	for i:=0;i<n;i++{
		if parents[i]==i{
			disconnectionCount+=1
		}
	}
	return disconnectionCount-1
}
