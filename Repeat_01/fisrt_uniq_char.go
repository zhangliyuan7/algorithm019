package Repeat_01

// 387
func firstUniqChar(s string) int {
	m:=make(map[string]int)
	for _,c:=range s{
		m[string(c)]+=1
	}
	for i,c:=range s{
		if m[string(c)]==1{
			return i
		}
	}
	return -1
}