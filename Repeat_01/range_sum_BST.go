package Repeat_01


//938
func rangeSumBST(root *TreeNode, low int, high int)(ans int ){
	var dg func(treeNode *TreeNode)
	dg = func(treeNode *TreeNode) {
		if treeNode!=nil{
			if treeNode.Val<=high&&treeNode.Val>=low{
				ans+=treeNode.Val
			}
			dg(treeNode.Left)
			dg(treeNode.Right)
		}
	}
	dg(root)
	return
}

func rangeSumBSTI(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

