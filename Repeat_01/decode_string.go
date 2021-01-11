package Repeat_01

import (
	"strconv"
)
//394 字符串解码
//3[a]4[ab]
//3[a4[ab]3[a]]=aababababaababababaabababab
func DecodeString(s string) string {
	stack:=[]string{}
	for i:=0;i<len(s);i++{
		if s[i]!=']'{
			stack=append(stack,string(s[i]))
		}else{
			tmpString:=""
			for stack[len(stack)-1]!="["{
				tmpString=stack[len(stack)-1]+tmpString
				stack=stack[:len(stack)-1]
			}
			//drop "["
			stack=stack[:len(stack)-1]
			numString :=""
			for len(stack)>0&&stack[len(stack)-1]>="0"&&stack[len(stack)-1]<="9"{
				numString=stack[len(stack)-1]+numString
				stack=stack[:len(stack)-1]
			}
			numInt,_:=strconv.Atoi(numString)
			tmp:=""
			for i:=0;i< numInt;i++{
				tmp+=tmpString
			}
			for i:=0;i<len(tmp);i++{
				stack=append(stack,string(tmp[i]))
			}
		}
	}
	r:=""
	for i:=0;i<len(stack);i++{
		r+=stack[i]
	}
	return r
}