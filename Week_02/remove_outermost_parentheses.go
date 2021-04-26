package Week_02

import "strings"

//栈为空时 入栈的元素索引，出栈后，栈为空的出栈元素索引 标记起来，最后统一删除即可
//speed 12ms  有点慢。。。
// bad，因为要重复处理删除 速度太慢
// 不如直接记录符合条件的string(c),而不是记录不符合条件的index
func removeOuterParentheses(S string) string {
	var s = []string{}
	var r = make(map[int ]bool)
	var result string
	for i,c:=range S{
		if string(c)=="("{
			if len(s)==0{
				r[i]=true
			}
			s=append(s,string(c))
		}
		if string(c)==")"{
			s=s[:len(s)-1]
			if len(s)==0{
				r[i]=true
			}
		}
	}
	for i,c:=range S{
		if _,ok:=r[i];!ok{
			result+=string(c)
		}
	}
	return result
}

//不记录要删除index，用strings.builder直接添加符合条件的元素到结果集
//good , 0ms
func removeOuterParenthesesBetter(S string) string {
	var s = []string{}
	var result strings.Builder
	for _,c:=range S{
		if string(c)=="("{
			if len(s)!=0{
				result.WriteString(string(c))
			}
			s=append(s,string(c))
		}
		if string(c)==")"{
			s=s[:len(s)-1]
			if len(s)!=0{
				result.WriteString(string(c))
			}
		}
	}

	return result.String()
}

// looks better than mine
// fast and small space
func removeOuterParentheses1(S string) string {
	if len(S)==0{
		return ""
	}
	stack := []byte{S[0]}
	start :=1
	var res strings.Builder
	for i:=1;i<len(S);i++{
		if len(stack)>0&&S[i]==')'&&stack[len(stack)-1]=='(' {
			stack = stack[:len(stack)-1]
			if len(stack)==0{
				res.WriteString(S[start:i])
				start = i+2
			}
		}else {
			stack = append(stack,S[i])
		}
	}
	return res.String()
}
