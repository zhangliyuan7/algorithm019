package w3

import "fmt"

// NO.22
var r = []string{}
//
func GenerateParenthesis(n int) []string {
	var s = ""
	//recursionGenerate1(0,n,s)
	recursionGenerate2(0,0,s,n)
	return r
}

//全排列
func recursionGenerate1(level int ,n int ,s string ){
	//terminator
	if level>n{
		//append()
		fmt.Println(s)
		return
	}
	//logic ,process
	//drill down
	recursionGenerate1(level+1,n,s+"(")
	recursionGenerate1(level+1,n,s+")")
	//recursion
}
// judge condition:left && right
func recursionGenerate2(left,right int ,s string ,n int ){
	//terminator
	if left==n&&right==n{
		r=append(r,s)
		return
	}
	//process
	//drill down
	if left<n{
		recursionGenerate2(left+1,right,s+"(",n)
	}
	if left>right{
		recursionGenerate2(left,right+1,s+")",n)
	}
	//recursion
}

func GenerateParenthesisA(n int) []string {
	var r =[]string{}
	var recursionFunc func(left,right int ,n int ,s string)
	recursionFunc=func(left,right int ,n int ,s string ){
		// 是 ==
		if left==n&&right==n{
			r=append(r,s)
			return
		}
		if left<n{
			recursionFunc(left+1,right,n,s+"(")
		}
		if right<left{
			recursionFunc(left,right+1,n,s+")")
		}
	}
	recursionFunc(0,0,n,"")
	return r
}








