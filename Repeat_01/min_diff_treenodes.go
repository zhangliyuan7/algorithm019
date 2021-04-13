package Repeat_01

import "math"

//530&&783
//midOrder
func minDiffInBST(root *TreeNode)( ans int) {
	pre:=-1
	ans=math.MaxInt32
	var dfs func(*TreeNode)
	dfs=func(treeNode *TreeNode){
		if treeNode==nil{
			return
		}
		dfs(root.Left)
		if pre!=-1&&ans>treeNode.Val-pre{
			ans=treeNode.Val-pre
		}
		pre=treeNode.Val
		dfs(treeNode.Right)
	}
	dfs(root)
	return ans
}
