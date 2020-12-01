package Week_06

func minMutation(start string, end string, bank []string) int {
	valid:=[]string{"A","T","C","G"}
	bankMap:=make(map[string]bool)
	for i:=range bank{
		if start==bank[i]{
			continue
		}
		bankMap[bank[i]]=true
	}
	if _,ok:=bankMap[end];!ok{
		return -1
	}
	visited:=make(map[string]bool) //只保留下一个正向层与反向层的节点
	visited[start]=true
	visited[end]=true

	s_q :=[]string{start}
	e_q:=[]string{end}
	count:=1

	for len(s_q)!=0{
		length:=len(s_q)  // 为了剔除队列中已经访问过的节点
		for _,k:=range s_q{
			for i,_:=range k{
				for _,r:=range valid{
					key:=k[0:i]+r+k[i+1:]
					if key==k{ // 自身continue
						continue
					}
					if _,ok:=bankMap[key];ok{
						if _,ok:=visited[key];ok{
							return count
						}
						visited[key]=true
						s_q=append(s_q,key)
					}
				}
			}
			delete(bankMap,k)  //删除已经看过的层
			delete(visited,k)
		}
		count+=1
		s_q=s_q[length:]
		if len(s_q)> len(e_q){
			e_q,s_q=s_q,e_q
		}
	}
	return -1
}
//bfs + double direct
func MinMutation(start string, end string, bank []string) int {
		valid:=[]string{"A","T","C","G"}
		bankMap:=make(map[string]bool)
		for i:=range bank{
		if start==bank[i]{
		continue
	}
		bankMap[bank[i]]=true
	}
		if _,ok:=bankMap[end];!ok{
		return -1
	}
		visited:=make(map[string]bool) //只保留下一个正向层与反向层的节点
		visited[start]=true
		visited[end]=true

		s_q :=[]string{start}
		e_q:=[]string{end}
		count:=1

		for len(s_q)!=0{
		length:=len(s_q)  // 为了剔除队列中已经访问过的节点
		tmpVisit:=[]string{} // 存储这层所有的新的准备访问的节点
		for _,k:=range s_q{
		for i,_:=range k{
		for _,r:=range valid{
		key:=k[0:i]+r+k[i+1:]
		if key==k{ // 自身continue
		continue
	}
		if _,ok:=visited[key];ok{
		return count
	}
		if _,ok:=bankMap[key];ok{
		tmpVisit=append(tmpVisit,key) // 防止同层转换出现多个同样的key，所以不在这里加入到visit
		s_q=append(s_q,key)
	}
	}
	}
		delete(visited, k)  //删除已经看过的层
	}
		for _,v:=range tmpVisit{ //因为正向和反向层 两层节点visit，不能直接覆盖
		visited[v]=true
	}
		count+=1
		s_q=s_q[length:]
		if len(s_q)> len(e_q){
		e_q,s_q=s_q,e_q
	}
	}
		return -1
	}

