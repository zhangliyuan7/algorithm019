package Week_08

// 387
func firstUniqCharL(s string) int {
	var slist =[26]int{}
	for _,c:=range s{
		slist[c-'a']+=1
	}
	for i,c:=range s{
		if slist[c-'a'] == 1{
			return i
		}
	}
	return -1
}
