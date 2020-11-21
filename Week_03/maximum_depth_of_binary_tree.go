package w3

// NO.104
type Node struct {
	val   int
	left  *Node
	right *Node
}

func MaxDepth(root *Node) int {
	var large = 0
	var recursionDepth func(*Node, int) int
	recursionDepth = func(root *Node, n int) int {
		if root == nil {
			return n
		}
		return max(recursionDepth(root.left, n+1), recursionDepth(root.right, n+1))
	}
	large = recursionDepth(root, 0)
	return large + 1
}

//代码精简 官方递归
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.left), maxDepth(root.right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//瞧着复杂不少。。。但是是原生态
func maxDepth3(root *Node) int {
	if root == nil {
		return 0
	}
	var max = 0
	var recursionDepth func(*Node, int)
	recursionDepth = func(root *Node, n int) {
		if root.right == nil && root.left == nil {
			if n > max {
				max = n
			}
			return
		}
		if root.left!=nil {
			recursionDepth(root.left, n+1)
		}
		if root.right!=nil {
			recursionDepth(root.right, n+1)
		}
	}
	recursionDepth(root, 0)
	return max + 1
}

//非递归 栈 dfs 遍历
func maxDepth4(root *Node)  {
	var s []*Node
	s = append(s, root)
	for len(s) != 0 {
		last := s[len(s)-1]
		s = s[:len(s)-1]
		root = last
		if root == nil {
			continue
		}
		s = append(s, root.right)
		s = append(s, root.left)
	}
}
