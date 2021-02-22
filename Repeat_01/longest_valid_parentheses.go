package Repeat_01

// 32
func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	r:=0
	stack := make([]int, len(s))
	stack=append(stack,-1)
	for i:=0;i<len(s);i++{
		if s[i]=='('{
			stack=append(stack,i)
		}else{
			stack:=stack[:len(stack)-1]
			if len(stack)==0{
				stack=append(stack,i)
			}else{
				if r>i-stack[len(stack)-1]{
					r=i-stack[len(stack)-1]
				}
			}
		}
	}
	return r
}

func longestValidParenthesesB(s string) int {
	lp,rp:=0,0
	ans:=0
	for i:=0;i<len(s);i++{
		if s[i]=='('{
			lp+=1
		}
		if s[i]==')'{
			rp+=1
		}
		if lp==rp{
			if ans<2*rp{
				ans=2*rp
			}
		}else if rp>lp{
			lp,rp=0,0
		}
	}
	lp,rp=0,0
	for i:=len(s)-1;i>=0;i--{
		if s[i]=='('{
			lp++
		}
		if s[i]==')'{
			rp++
		}
		if lp==rp{
			if ans<2*lp{
				ans=2*lp
			}
		}else if lp>rp{
			lp,rp=0,0
		}
	}
	return ans
}