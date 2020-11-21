package w3

import "math"

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

func isValidBST(root *TreeNode) bool {
	if root==nil{
		return true
	}
	var s []*TreeNode
	var last =math.MinInt64
	for root!=nil || len(s)!=0{
		if root!=nil{
			s=append(s,root)
			root=root.Left
			continue
		}
		if len(s)!=0{
			root=s[len(s)-1]
			if last>=root.Val{
				return false
			}
			last=root.Val
			s=s[:len(s)-1]
			root=root.Right
		}
	}
	return true
}