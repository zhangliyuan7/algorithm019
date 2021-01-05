package Week_10

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后序遍历 double stack
// 145
func postOrder(node *TreeNode) []int {
	r := []int{}
	if node == nil {
		return r
	}
	s1 := []*TreeNode{}
	s2 := []*TreeNode{}
	s1 = append(s1, node)
	for len(s1) != 0 {
		top := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]

		s2 = append(s2, top)
		if top.Left != nil {
			s1 = append(s1, top.Left)
		}
		if top.Right != nil {
			s1 = append(s1, top.Right)
		}
	}
	for i := len(s2) - 1; i >= 0; i-- {
		r = append(r, s2[i].Val)
	}
	return r
}
