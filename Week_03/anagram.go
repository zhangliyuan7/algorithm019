package w3

import "sort"

//fast
func isAnagram(s string, t string) bool {
	var list1,list2 [26]int
	for _,c:=range s{
		list1[c-97]+=1
	}
	for _,c:=range t{
		list2[c-97]+=1
	}
	if list1==list2{
		return true
	}
	return false
}
//slow
func isAnagram2(s string,t string )bool{
	sb:=[]byte(s)
	tb:=[]byte(t)
	sort.Slice(sb, func(i, j int) bool {
		if sb[i]<sb[j]{
			return true
		}
		return false
	})
	sort.Slice(tb, func(i, j int) bool {
		if tb[i]<tb[j]{
			return true
		}
		return false
	})
	if string(sb)==string(tb){
		return true
	}
	return false
}
