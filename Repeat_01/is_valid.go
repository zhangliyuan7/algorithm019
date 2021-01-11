package Repeat_01

// 20 有效的括号
func isValid(s string) bool {
	pars:=make(map[string]string)
	pars[")"]="("
	pars["]"]="["
	pars["}"]="{"
	stack:=[]string{}
	for i:=0;i<len(s);i++{
		if s[i]== '('||s[i]=='['||s[i]=='{'{
			stack=append(stack,string(s[i]))
			continue
		}
		if v,ok:=pars[string(s[i])];ok && len(stack)!=0 && v==stack[len(stack)-1]{
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	return true
}