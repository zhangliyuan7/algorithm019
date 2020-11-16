package w2

//hash baoli : rubbish 20ms
func isAnagram(s string, t string) bool {
	if len(s)!=len(t){
		return false
	}
	var h = make(map[string]int)
	for _,c:=range s{
		if v,ok:=h[string(c)];ok{
			h[string(c)]=v+1
		}else{
			h[string(c)]=1
		}
	}
	for _,c2:=range t{
		if _,ok:=h[string(c2)];!ok{
			return false
		}
		h[string(c2)]-=1
		if h[string(c2)]==0{delete(h,string(c2))}
	}
	if len(h)==0{
		return true
	}
	return false
}

// bull shit : 0ms
func IsAnagram(s string,t string )bool{
	if len(s)!=len(t){return false }
	var sArr = [26]int{}
	var tArr = [26]int{}
	for i:=0;i<len(s);i++{
		sArr[s[i]-'a']++
		tArr[t[i]-'a']++
	}
	if sArr==tArr{
		return true
	}
	return false
}