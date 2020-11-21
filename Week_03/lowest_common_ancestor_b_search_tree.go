package w3

func lowestCommonAncestorBinarySearchTree(root, p, q *TreeNode) *TreeNode {
	if root==nil||root==p||root==q{
		return root
	}
	s,l:=p.Val,q.Val
	if s>l{
		s,l=l,s
	}
	return getAncestor(root,s,l)
}

func getAncestor(root *TreeNode,s,l int ) *TreeNode {
	if root==nil||root.Val==s||root.Val==l||(s<root.Val&&root.Val<l){
		return root
	}
	if root.Val<s{
		return getAncestor(root.Right,s,l)
	}
	if root.Val>l{
		return getAncestor(root.Left,s,l)
	}
	return nil
}