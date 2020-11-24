package Week_04

func levelOrder(root *TreeNode) [][]int {
	var r [][]int
	if root==nil{
		return r
	}
	var q []*TreeNode
	q=append(q,root)
	for len(q)!=0{
		r=append(r,vals(q))
		tmp_q :=[]*TreeNode{}
		for len(q)!=0{
			root:=q[0]
			q=q[1:]
			if root.Left!=nil{
				tmp_q=append(tmp_q,root.Left)
			}
			if root.Right!=nil{
				tmp_q=append(tmp_q,root.Right)
			}
		}
		q=tmp_q
	}
	return r
}
func vals(q []*TreeNode)[]int{
	var r []int
	for i:=0;i<len(q);i++{
		r=append(r,q[i].Val)
	}
	return r
}