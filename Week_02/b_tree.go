package Week_02

import "fmt"

type Node struct{
	Val int
	L *Node
	R *Node
}

//中序遍历 左-根-右 递归办法
func inorderTraversal(root *Node) []int {
	var r []int
	return recursionInor(root,r)
}

func recursionInor(root *Node,r []int)[]int{
	if root!=nil {
		r = recursionInor(root.L, r)
		r = append(r, root.Val)
		r = recursionInor(root.R, r)
	}
	return r
}

// 先序递归 root-L-R
func preTraversal(root *Node) []int {
	var r []int
	return recursionPre(root,r)
}

func recursionPre(root *Node,r []int) []int {
	if root!=nil{
		r=append(r,root.Val)
		r= recursionPre(root.L,r)
		r=recursionPre(root.R,r)
	}
	return r
}

//后序递归 L-R-root

func backTraversal(root *Node) []int {
	var r []int
	return recursionBack(root,r)
}
func recursionBack(root *Node,r []int) []int{
	if root!=nil{
		r =recursionBack(root,r)
		r = recursionBack(root,r)
		r=append(r,root.Val)
	}
	return r
}

// 非递归
// 中序 左中右
func inorderTraversal1(root *Node) []int {
	var r []int
	var s []*Node
	var p = root
	for p!=nil||len(s)!=0{
		for p!=nil {
			s=append(s,p)
			p=p.L
		}
		if len(s)!=0{
			p=s[len(s)-1]
			s=s[:len(s)-1]
			r=append(r,p.Val)
			p=p.R
		}
	}
	return r
}

//先序 中左右
func preTr(root *Node)[]int{
	var r []int
	var s []*Node
	var p = root
	for p!=nil||len(s)!=0{
		for p!=nil {
			r = append(r,p.Val)
			s = append(s, p)
			p = p.L
		}
		if len(s)!=0{
			p=s[len(s)-1]
			s=s[:len(s)-1]
			p=p.R
		}
	}
	return r
}

//后序 左右中
func bcktr(root *Node)[]int{
	return nil
}


func xianxudiguiA(root *Node){
	for root!=nil{
		fmt.Println(root.Val)
		xianxudiguiA(root.L)
		xianxudiguiA(root.R)
	}
}
func zhongxudiguiA(root *Node){
	for root!=nil{
		xianxudiguiA(root.L)
		fmt.Println(root.Val)
		xianxudiguiA(root.R)
	}
}
func houxudiguiA(root *Node){
	for root!=nil{
		xianxudiguiA(root.L)
		xianxudiguiA(root.R)
		fmt.Println(root.Val)
	}
}

func xianxuA(root *Node){
	var s []*Node
	for root!=nil||len(s)!=0{
		for root!=nil{
			fmt.Println(root.Val)
			s=append(s,root)
			root=root.L
		}
		if len(s)!=0{
			root=s[len(s)-1]
			s=s[:len(s)-1]
			root=root.R
		}
	}
}

func zhongxuA(root *Node){
	var s []*Node
	for root!=nil||len(s)!=0{
		for root!=nil{
			s=append(s,root)
			root=root.L
		}
		if len(s)!=0{
			root=s[len(s)-1]
			fmt.Println(root.Val)
			s=s[:len(s)-1]
			root=root.R
		}
	}
}

func houxuA(root *Node){

}