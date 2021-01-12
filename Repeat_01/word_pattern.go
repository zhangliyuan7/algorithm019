package Repeat_01

import ("strings")
// 290
func wordPattern(pattern string, s string) bool {
	m:=make(map[string]string)
	m2:=make(map[string]string)
	words:=strings.Split(s," ")
	if len(words)!=len(pattern){ return false }
	for i,c:=range pattern{
		if v,ok:=m[string(c)];ok{
			if v==words[i]{
				continue
			}else{
				return false
			}
		}
		if _,ok:=m2[words[i]];ok{
			return false
		}
		m[string(c)]=words[i]
		m2[words[i]]=string(c)
	}
	return true
}