package Week_04

import (
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 遍历，广度优先
func countNodesBfs(root *TreeNode) int {
	if root==nil{
		return 0
	}
	var n int
	var queue []*TreeNode
	queue=append(queue,root)
	for len(queue)!=0{
			node:=queue[0]
			queue=queue[1:]
			n+=1
			if node.Left!=nil{
				queue=append(queue,node.Left)
			}
			if node.Right!=nil{
				queue=append(queue,node.Right)
			}
	}
	return n
}
//深度优先
func countdfs(root *TreeNode)int{
	if root ==nil{
		return 0
	}
	var s []*TreeNode
	s=append(s,root)
	var n int
	for len(s)!=0{
		root:=s[len(s)-1]
		s=s[:len(s)-1]
		n+=1
		if root.Right!=nil{
			s=append(s, root.Right)
		}
		if root.Left!=nil{
			s=append(s, root.Left)
		}
	}
	return n
}

//中序/先序遍历 比较快
func countNodesPre(root *TreeNode)int{
	var s []*TreeNode
	var n int
	for root!=nil||len(s)!=0{
		for root!=nil{
			s=append(s,root)
			root=root.Left
		}
		if len(s)!=0{
			root=s[len(s)-1]
			n+=1
			s=s[:len(s)-1]
			//zhongxu
			root=root.Right
		}
	}
	return n
}

//广度优先 找层级 统计层数 ，自己想的果然比较麻烦。。。
func countNodes(root *TreeNode) int {
	if root==nil{
		return 0
	}
	if root.Left==nil&&root.Right==nil{
		return 1
	}
	var n int
	var queue []*TreeNode
	queue=append(queue,root)
	for len(queue)!=0{
		var  tmp []*TreeNode
		for len(queue)!=0{
			node:=queue[0]
			queue=queue[1:]
			if node.Left!=nil{
				tmp=append(tmp,node.Left)
			}
			if node.Right!=nil{
				tmp=append(tmp,node.Right)
			}
		}
		queue=tmp
		n+=1
		if len(queue)!=int(math.Pow(2,float64(n))){
			return Count(n-1)+len(queue)
		}
	}
	return Count(n)
}
func Count(n int )int{
	var sum int
	for i:=0;i<n;i++{
		sum+=int(math.Pow(2,float64(i)))
	}
	return sum
}

//简洁 暴力遍历
func digui(root *TreeNode)int{
	if root==nil{
		return 0
	}
	return digui(root.Left)+digui(root.Right)+1
}

func erfen(root *TreeNode)int{
	if root==nil{
		return 0
	}
	left:=level(root.Left)
	right:=level(root.Right)
	//子树count个数应该是(1<<left)-1
	//但是要加上root节点 ，所以是(1<<left)
	if left==right{
		return erfen(root.Right)+(1<<left)
	}else{
		return erfen(root.Left)+(1<<right)
	}
}

func level(root *TreeNode)int{
	var n int
	for root!=nil{
		n+=1
		root=root.Left
	}
	return n
}