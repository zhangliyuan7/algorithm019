package Repeat_01

//897
func increasingBST(root *TreeNode) *TreeNode {
	newTree:=new(TreeNode)
	base:=newTree
	var midOrderBST func(*TreeNode)
	midOrderBST=func(root *TreeNode){
		if root==nil{
			return
		}
		midOrderBST(root.Left)
		newTree.Right=root
		root.Left=nil
		newTree=root
		midOrderBST(root.Right)
	}
	midOrderBST(root)
	return base.Right
}
