package Repeat_01

// 3 //double pointers
func LengthOfLongestSubstring(s string )int {
	m:=make(map[string]int)
	left:=0
	right:=0
	maxLength:=0
	for i,c:=range s {
		if v,ok:=m[string(c)];ok&&v>=left{
			if right-left>maxLength{
				maxLength=right-left
			}
			left=v+1
			right+=1
			m[string(c)]=i
			continue
		}
		right+=1
		m[string(c)]=i
	}
	if right-left>maxLength{
		return right-left
	}
	return maxLength
}