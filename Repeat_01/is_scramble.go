package Repeat_01

// 87
var memo map[string]bool
func isScramble(s1 string, s2 string) bool {
	memo = make(map[string]bool)
	return dfsIsScramble(s1,s2)
}
func dfsIsScramble(s1 ,s2 string )bool{
	key :=s1+s2
	if value, ok := memo[key]; ok {
		return value
	}
	if s1==""||s2==""||len(s1)!=len(s2){
		memo[key]=false
		return false
	}
	if s1==s2{
		memo[key]=true
		return true
	}
	//判断包含的字母是否相同
	if hashCode(s1)!=hashCode(s2){
		memo[key]=false
		return false
	}
	// m:=make(map[string]int)
	// for i:=range s1{
	// 	m[string(s1[i])]+=1
	// 	m[string(s2[i])]-=1
	// }
	// for _,v:=range m{
	// 	if v!=0{
	// 		memo[key]=false
	// 		return false
	// 	}
	// }
	ln:=len(s1)
	for i:=1;i<ln;i++ {
		if dfsIsScramble(s1[:i], s2[:i]) && dfsIsScramble(s1[i:], s2[i:]) {
			memo[key]=true
			return true
		}
		//dfs(s1[:i], s2[l1-i:]) && dfs(s1[i:], s2[:l1-i])
		if dfsIsScramble(s1[:i], s2[ln-i:]) && dfsIsScramble(s1[i:], s2[:ln-i]) {
			memo[key]=true
			return true
		}
	}
	memo[key]=false
	return false
}

func hashCode(s string )int{
	x:=0
	for i:=range s{
		x+=1<<(s[i]-'a')
	}
	return x
}