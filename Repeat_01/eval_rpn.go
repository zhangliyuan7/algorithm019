package Repeat_01

import "strconv"

//150
func evalRPN(tokens []string) int {
	ln:=len(tokens)
	stack:=[]int{}
	for i:=0;i<ln;i++{
		ele:=tokens[i]
		if ele!="+"&&ele!="-"&&ele!="*"&&ele!="/"{
			n,_:=strconv.Atoi(ele)
			stack=append(stack,n)
		}else{
			n2:=stack[len(stack)-1]
			n1:=stack[len(stack)-2]
			stack=stack[:len(stack)-2]
			switch ele{
			case "+":
				stack=append(stack,n1+n2)
			case "-":
				stack=append(stack,n1-n2)
			case "*":
				stack=append(stack,n1*n2)
			case "/":
				stack=append(stack,n1/n2)
			}
		}
	}
	return stack[0]
}
