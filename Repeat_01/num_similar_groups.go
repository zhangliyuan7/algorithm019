package Repeat_01

// 839
func numSimilarGroups(strs []string) int {
	l:=len(strs)
	parents:=make([]int,l)
	strMap:=make(map[string]int)
	for i:=range strs{
		strMap[strs[i]]=i
	}
	var find func(x int  )int
	find=func(x int )int{
		if parents[x]!=x{
			parents[x]=find(parents[x])
		}
		return parents[x]
	}
	var union func(x,y int )
	union=func(x,y int ){
		px:=find(x)
		py:=find(y)
		if px == py {
			return
		}
		parents[py]=px
		parents[y]=px
	}
	for i:=0;i<l;i++{
		parents[i]=i
	}
	for i:=0;i<len(strs);i++ {
		for j := i+1; j < len(strs); j++ {
			if isSimilar(strs[i], strs[j]) {
				union(i, j)
			}
		}
	}
	count:=0
	for i:=range parents{
		if parents[i]==i{
			count+=1
		}
	}
	return count
}

//func SearchSimilarWord(str string,strMap map[string]int )map[string]int{
//	newMap:=make(map[string]int)
//	for i:=0;i<len(str);i++{
//		tmp:=str
//		for j:=0;j<len(tmp);j++{
//			byteTmp:=[]byte(tmp)
//			byteTmp[i],byteTmp[j]=byteTmp[j],byteTmp[i]
//				if v,ok:=strMap[string(byteTmp)];ok{
//					newMap[string(byteTmp)]=v
//				}
//		}
//	}
//	return newMap
//}

func isSimilar(s, t string) bool {
	diff := 0
	for i := range s {
		if s[i] != t[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	return true
}
