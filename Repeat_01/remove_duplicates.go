package Repeat_01

//1047 stack
func removeDuplicates(S string) string {
	s:=[]byte{}
	for i:=range S{
		if len(s)==0 || s[len(s)-1]!=S[i]{
			s=append(s,S[i])
		}else {
			s=s[:len(s)-1]
		}
	}
	return string(s)
}
